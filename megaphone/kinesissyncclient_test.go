package megaphone

import (
	"testing"
	"github.com/stretchr/testify/require"
	. "github.com/petergtz/pegomock"
	"github.com/redbubble/megaphone-client-golang/megaphone/mock"
	"reflect"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"errors"
)

func TestKinesisSyncClient(t *testing.T) {

	var required *require.Assertions
	var awsKinesisClient *mock.MockKinesisAPI
	var kinesisSyncClient *KinesisSyncClient

	setup := func(t *testing.T) {
		RegisterMockTestingT(t)

		awsKinesisClient = mock.NewMockKinesisAPI()
		kinesisSyncClient = &KinesisSyncClient{
			kinesisClient: awsKinesisClient,
			clientConfig: KinesisClientConfig{
				Config:    Config{"test-client"},
				DeployEnv: "test-env",
			},
		}
		required = require.New(t)
	}

	t.Run("Publish()", func(t *testing.T) {
		t.Run("calls kinesisClient.PutRecord", func(t *testing.T) {
			setup(t)

			payload := `{
						  "id": 7,
						  "foo": "bar"
						}`
			err := kinesisSyncClient.Publish("test-topic", "test-subtopic", "http://schema.org/test.json", "42", []byte(payload))

			required.NoError(err)

			input := awsKinesisClient.VerifyWasCalledOnce().PutRecord(anyPutRecordInput()).GetCapturedArguments()

			expectedPayload := `{
								  "origin": "test-client",
								  "topic": "test-topic",
								  "subtopic": "test-subtopic",
								  "schema": "http://schema.org/test.json",
								  "partitionKey": "42",
								  "data": {
									"foo": "bar",
									"id": 7
								  }
								}`
			required.Equal("42", *input.PartitionKey)
			required.Equal("megaphone-streams-test-env-test-topic", *input.StreamName)
			required.JSONEq(expectedPayload, string(input.Data))
		})

		t.Run("returns error if invalid payload", func(t *testing.T) {
			setup(t)

			payload := `{
						  "foo"
						}`
			err := kinesisSyncClient.Publish("test-topic", "test-subtopic", "http://schema.org/test.json", "42", []byte(payload))

			required.Error(err)
			awsKinesisClient.VerifyWasCalled(Never()).PutRecord(anyPutRecordInput())
		})

		t.Run("returns error if kinesis client returns an error", func(t *testing.T) {
			setup(t)

			payload := `{
						  "id": 7,
						  "foo": "bar"
						}`
			When(awsKinesisClient.PutRecord(anyPutRecordInput())).ThenReturn(nil, errors.New("aws-error"))
			err := kinesisSyncClient.Publish("test-topic", "test-subtopic", "http://schema.org/test.json", "42", []byte(payload))

			required.Error(err)
			awsKinesisClient.VerifyWasCalledOnce().PutRecord(anyPutRecordInput())
		})
	})
}

func anyPutRecordInput() *kinesis.PutRecordInput {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((*kinesis.PutRecordInput)(nil))))
	return nil
}
