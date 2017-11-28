package logger

import (
	"fmt"
	"os"
)

type fileLogger struct {
}

// NewFileLogger returns a configured file logger to publish events to a local file,
// mainly for development purposes.
func NewFileLogger() (*fileLogger, error) {
	return &fileLogger{}, nil
}

func (fileLogger *fileLogger) Post(topic string, message interface{}) (err error) {
	f, err := os.OpenFile(getFileName(topic), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	if _, err = f.WriteString(message.(string)); err != nil {
		return
	}

	return
}

func getFileName(topic string) string {
	return fmt.Sprintf("%s.stream", topic)
}
