package logger_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/redbubble/megaphone-client-golang/megaphone/logger"
	"github.com/stretchr/testify/assert"
)

func TestFileLogger(t *testing.T) {

	getFileName := func(topic string) string {
		return fmt.Sprintf("%s.stream", topic)
	}

	t.Run("Post()", func(t *testing.T) {

		t.Run("it writes in the file", func(t *testing.T) {
			topic := "work-updates"
			payload := "test_payload"
			fileName := getFileName(topic)
			os.Remove(fileName)

			fileLogger, _ := logger.NewFileLogger()

			err := fileLogger.Post(topic, payload)
			assert.Nil(t, err)

			data, err := ioutil.ReadFile(fileName)
			assert.Nil(t, err)
			assert.Equal(t, payload, string(data))
		})
	})
}
