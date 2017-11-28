// Package megaphone provides a client for Megaphone, Redbubble's event
// broadcasting system for inter-service communication.
package megaphone

import (
	"os"
	"strconv"

	"github.com/redbubble/megaphone-client-golang/megaphone/logger"
)

// Config holds the configuration for a Client.
type Config struct {
	Origin string
	Host   string
	Port   int
}

// Client provides a standard interface to publish events to Megaphone.
// Conventionally, a FileLogger is used if no configuration for
// a FluentLogger is provided, allowing for easy local development
// in absence of a local Megaphone Fluentd container.
type Client struct {
	config Config
	logger logger.Logger
}

// NewClient returns a configured Megaphone Client.
func NewClient(config Config) (client *Client, err error) {
	client = &Client{}
	client.config = config

	if client.config.Host == "" {
		client.config.Host = os.Getenv("MEGAPHONE_FLUENT_HOST")
	}

	port := os.Getenv("MEGAPHONE_FLUENT_PORT")
	if port != "" {
		client.config.Port, err = strconv.Atoi(port)
		if err != nil {
			return client, NewConfigErrorWithField(err, "port")
		}
	}

	client.logger, err = logger.NewLogger(config.Host, config.Port)
	if err != nil {
		return client, NewConfigError(err)
	}

	return
}

// Publish sends an event to Megaphone, or to a local file depending on the Client configuration.
func (c *Client) Publish(origin, topic, subtopic, schema, partitionKey string, payload []byte) (err error) {
	event, err := newEvent(origin, topic, subtopic, schema, partitionKey, payload)
	if err != nil {
		return NewPayloadError(err, string(payload))
	}

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
