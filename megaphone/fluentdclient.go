// Package megaphone provides a client for Megaphone, Redbubble's event
// broadcasting system for inter-service communication.
package megaphone

import (
	"os"
	"strconv"

	"github.com/redbubble/megaphone-client-golang/megaphone/logger"
)

// FluentdConfig holds the configuration for a FluentdClient.
type FluentdConfig struct {
	Origin string
	Host   string
	Port   int
}

// Client provides a standard interface to publish events to Megaphone.
// Conventionally, a FileLogger is used if no configuration for
// a FluentLogger is provided, allowing for easy local development
// in absence of a local Megaphone Fluentd container.
type FluentdClient struct {
	logger logger.Logger
	FluentdConfig
}

func newConfig(origin, host string, port int) (FluentdConfig, error) {
	config := FluentdConfig{
		Origin: origin,
		Host:   host,
		Port:   port,
	}

	if config.Host == "" {
		config.Host = os.Getenv("MEGAPHONE_FLUENT_HOST")
	}

	if port == 0 {
		sPort := os.Getenv("MEGAPHONE_FLUENT_PORT")
		if sPort != "" {
			var err error
			config.Port, err = strconv.Atoi(sPort)
			if err != nil {
				return config, NewConfigErrorWithField(err, "port")
			}
		}
	}

	return config, nil
}

// NewFluentdClient returns a configured Megaphone FluentdClient.
func NewFluentdClient(origin, host string, port int) (c Publisher, err error) {
	client := &FluentdClient{}
	client.FluentdConfig, err = newConfig(origin, host, port)
	if err != nil {
		return client, err
	}

	var errLogger error
	client.logger, errLogger = logger.NewLogger(client.Host, client.Port)
	if errLogger != nil {
		return client, NewConfigError(errLogger)
	}

	c = client
	return c, nil
}

// Publish sends an event to Megaphone, or to a local file depending on the Client configuration.
func (c *FluentdClient) Publish(topic, subtopic, schema, partitionKey string, payload []byte) (err error) {
	event, err := newEvent(topic, subtopic, schema, partitionKey, payload)
	if err != nil {
		return NewPayloadError(err, string(payload))
	}
	event.Origin = c.Origin

	eventJSON, err := event.toJSON()
	if err != nil {
		return NewPublicationError(err, eventJSON)
	}

	err = c.logger.Post(topic, eventJSON)
	if err != nil {
		return NewPublicationError(err, eventJSON)
	}
	return nil
}
