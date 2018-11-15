package megaphone

// Publisher provides means to publish an event.
type Publisher interface {
	Publish(topic, subtopic, schema, partitionKey string, payload []byte) (err error)
	PublishRawMessage(streamName string, partitionKey string, messageBytes []byte) (err error)
}
