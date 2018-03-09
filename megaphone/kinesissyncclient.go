package megaphone

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

type KinesisSyncClient struct {
	kinesisClient kinesisiface.KinesisAPI
	clientConfig  KinesisClientConfig
}

func (c *KinesisSyncClient) Publish(topic, subtopic, schema, partitionKey string, payload []byte) error {
	event, err := newEvent(topic, subtopic, schema, partitionKey, payload)
	if err != nil {
		return err
	}
	event.Origin = c.clientConfig.Origin
	bytes, err := event.toJSONBytes()
	if err != nil {
		return err
	}
	input := &kinesis.PutRecordInput{
		Data:         bytes,
		PartitionKey: &partitionKey,
		StreamName:   aws.String(event.streamName(c.clientConfig.DeployEnv)),
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
