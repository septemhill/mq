package sqs

import (
	"context"

	"github.com/t/mq"
)

type SQSConsumerService struct {
	repo mq.MQConsumerRepository
}

func (s *SQSConsumerService) Fetch(ctx context.Context) ([]byte, error) {
	return s.repo.Fetch(ctx)
}

func NewSQSConsumerService(repo mq.MQConsumerRepository) *SQSConsumerService {
	return &SQSConsumerService{repo: repo}
}
