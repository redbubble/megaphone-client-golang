package megaphone

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/redbubble/megaphone-client-golang/megaphone/kinesisclient"
)

type KinesisSyncClient struct {
	kinesisClient kinesisiface.KinesisAPI
	config        kinesisclient.Config
}

func NewKinesisSyncClient(config kinesisclient.Config) (*KinesisSyncClient, error) {
	client, err := kinesisclient.ProvideKinesisClient(config)
	if err != nil {
		return nil, err
	}
	return &KinesisSyncClient{
		config:        config,
		kinesisClient: client,
	}, nil
}

func (c *KinesisSyncClient) Publish(topic, subtopic, schema, partitionKey string, payload []byte) error {
	event, err := newEvent(topic, subtopic, schema, partitionKey, payload)
	if err != nil {
		return err
	}
	event.Origin = c.config.Origin
	bytes, err := event.toJSONBytes()
	if err != nil {
		return err
	}
	input := &kinesis.PutRecordInput{
		Data:         bytes,
		PartitionKey: &partitionKey,
		StreamName:   aws.String(event.streamName(c.config.DeployEnv)),
	}
	err = input.Validate()
	if err != nil {
		return err
	}
	_, err = c.kinesisClient.PutRecord(input)
	if err != nil {
		return err
	}
	return nil
}
