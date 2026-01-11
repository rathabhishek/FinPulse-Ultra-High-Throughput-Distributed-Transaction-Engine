package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// SendToKinesis reflects your experience with real-time data streaming at Amazon [cite: 32]
func SendToKinesis(data []byte, streamName string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc := kinesis.New(sess)

	_, err := svc.PutRecord(&kinesis.PutRecordInput{
		Data:         data,
		PartitionKey: aws.String("partition-1"),
		StreamName:   aws.String(streamName),
	})
	return err
}
