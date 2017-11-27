package megaphone

import (
	"os"
	"strconv"

	"github.com/redbubble/megaphone-client-golang/megaphone/logger"
)

type Config struct {
	Origin string
	Host   string
	Port   int
}

type Client struct {
	config Config
	logger logger.Logger
}

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
