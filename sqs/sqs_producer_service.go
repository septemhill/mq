package sqs

import (
	"context"

	"github.com/t/mq"
)

type SQSProducerService struct {
	repo mq.MQProducerRepository
}

func (s *SQSProducerService) Push(ctx context.Context, b []byte) error {
	return s.repo.Push(ctx, b)
}

func (s *SQSProducerService) BatchPush(ctx context.Context, bs [][]byte) error {
	return s.repo.BatchPush(ctx, bs)
}

var _ mq.MQProducerService = (*SQSProducerService)(nil)

func NewSQSProducerService(repo mq.MQProducerRepository) *SQSProducerService {
	return &SQSProducerService{repo: repo}
}
