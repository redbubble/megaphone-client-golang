package kinesisclient

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func Provide(config Config) (*kinesis.Kinesis, error) {
	if config.HostedOnAWS {
		sess := session.Must(session.NewSession(&aws.Config{
			LogLevel: aws.LogLevel(aws.LogOff),
		}))
		region, regionErr := ec2metadata.New(sess).Region()
		if regionErr != nil {
			return nil, regionErr
		}
		creds := ec2rolecreds.NewCredentials(sess)
		return kinesis.New(sess, aws.NewConfig().WithCredentials(creds).WithRegion(region)), nil
	} else {
		sess := session.Must(session.NewSession(&aws.Config{
			LogLevel: aws.LogLevel(aws.LogOff),
		}))
		return kinesis.New(sess, aws.NewConfig().
			WithEndpoint(os.Getenv("MEGAPHONE_KINESIS_TEST_ENDPOINT")).
			WithDisableSSL(true).
			WithRegion(endpoints.UsEast1RegionID)), nil
	}
}
