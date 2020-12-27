package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/t/mq"
)

type SQSProducerRepository struct {
	bucketURL string
	sqs       *sqs.SQS
}

func (repo *SQSProducerRepository) Push(ctx context.Context, b []byte) error {
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    aws.String(repo.bucketURL),
	}
	_, err := repo.sqs.SendMessageWithContext(ctx, input)
	return err
}

func (repo *SQSProducerRepository) BatchPush(ctx context.Context, bs [][]byte) error {
	var entries []*sqs.SendMessageBatchRequestEntry
	for i := 0; i < len(bs); i++ {
		var e sqs.SendMessageBatchRequestEntry
		e.MessageBody = aws.String(string(bs[i]))
		entries = append(entries, &e)
	}

	input := &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(repo.bucketURL),
	}

	_, err := repo.sqs.SendMessageBatchWithContext(ctx, input)
	return err
}

var _ mq.MQProducerRepository = (*SQSProducerRepository)(nil)

func NewSQSProducerRepository(bucketURL string) *SQSProducerRepository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))

	return &SQSProducerRepository{
		bucketURL: bucketURL,
		sqs:       sqs.New(sess),
	}
}
