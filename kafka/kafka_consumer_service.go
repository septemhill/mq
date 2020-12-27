package kafka

import (
	"context"

	"github.com/t/mq"
)

type KafkaConsumerService struct {
	repo mq.MQConsumerRepository
}

func (s *KafkaConsumerService) Fetch(ctx context.Context) ([]byte, error) {
	return s.repo.Fetch(ctx)
}

var _ mq.MQConsumerRepository = (*KafkaConsumerService)(nil)

func NewKafkaConsumerService(repo mq.MQConsumerRepository) *KafkaConsumerService {
	return &KafkaConsumerService{repo: repo}
}
