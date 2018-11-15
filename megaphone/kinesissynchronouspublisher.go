package megaphone

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/redbubble/megaphone-client-golang/megaphone/kinesisclient"
)

type KinesisSynchronousPublisher struct {
	kinesisClient kinesisiface.KinesisAPI
	config        kinesisclient.Config
}

func NewKinesisSynchronousPublisher(config kinesisclient.Config) (*KinesisSynchronousPublisher, error) {
	client, err := kinesisclient.Provide(config)
	if err != nil {
		return nil, err
	}
	return &KinesisSynchronousPublisher{
		config:        config,
		kinesisClient: client,
	}, nil
}

func (c *KinesisSynchronousPublisher) Publish(topic, subtopic, schema, partitionKey string, payload []byte) error {
	event, err := newEvent(topic, subtopic, schema, partitionKey, payload)
	if err != nil {
		return err
	}
	event.Origin = c.config.Origin
	bytes, err := event.toJSONBytes()
	if err != nil {
		return err
	}
	err = c.PublishRawMessage(event.streamName(c.config.DeployEnv), partitionKey, bytes)
	if err != nil {
		return err
	}
	return nil
}

func (c *KinesisSynchronousPublisher) PublishRawMessage(streamName string, partitionKey string, messageBytes []byte) error {
	input := &kinesis.PutRecordInput{
		Data:         messageBytes,
		PartitionKey: &partitionKey,
		StreamName:   aws.String(streamName),
	}
	err := input.Validate()
	if err != nil {
		return err
	}
	_, err = c.kinesisClient.PutRecord(input)
	if err != nil {
		return err
	}
	return nil
}
