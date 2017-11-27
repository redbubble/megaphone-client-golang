package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Run("NewLogger()", func(t *testing.T) {
		t.Run("it returns a Fluent logger", func(t *testing.T) {
			host := os.Getenv("MEGAPHONE_FLUENT_HOST")
			if host == "" {
				host = "localhost"
			}
			port := 24224
			logger, err := NewLogger(host, port)
			assert.Nil(t, err)
			_, ok := logger.(*fluentLogger)
			assert.Equal(t, true, ok)
		})

		t.Run("it returns a file logger", func(t *testing.T) {
			host := ""
			port := 0
			logger, err := NewLogger(host, port)
			assert.Nil(t, err)
			_, ok := logger.(*fileLogger)
			assert.Equal(t, true, ok)
		})
	})
}
