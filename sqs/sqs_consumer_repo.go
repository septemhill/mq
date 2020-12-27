package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSConsumerRepository struct {
	bucketURL string
	sqs       *sqs.SQS
}

func (repo *SQSConsumerRepository) Fetch(ctx context.Context) ([]byte, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:          aws.String(repo.bucketURL),
		WaitTimeSeconds:   aws.Int64(20),
		VisibilityTimeout: aws.Int64(900),
	}

	rsp, err := repo.sqs.ReceiveMessage(input)
	if err != nil {
		return nil, err
	}

	return []byte(*rsp.Messages[0].Body), nil
}

func NewSQSConsumerRepository(bucketURL string) *SQSConsumerRepository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))

	return &SQSConsumerRepository{
		bucketURL: bucketURL,
		sqs:       sqs.New(sess),
	}
}
