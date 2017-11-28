package logger

import (
	"encoding/json"
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
)

type fluentLogger struct {
	Fluent *fluent.Fluent
}

// NewFluentLogger returns a Logger that can publish events to Megaphone
// via a Megaphone Fluentd container.
func NewFluentLogger(fluentHost string, fluentPort int) (logger *fluentLogger, err error) {
	logger = &fluentLogger{}
	logger.Fluent, err = fluent.New(fluent.Config{
		FluentHost: fluentHost,
		FluentPort: fluentPort,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (logger *fluentLogger) Post(topic string, message interface{}) (err error) {
	var payloadJSON map[string]interface{}
	err = json.Unmarshal([]byte(message.(string)), &payloadJSON)
	if err != nil {
		return err
	}

	err = logger.Fluent.Post(topic, payloadJSON)
	if err != nil {
		return
	}
	return
}
