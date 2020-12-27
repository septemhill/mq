package kafka

import (
	"context"

	"github.com/t/mq"
)

type KafkaProducerService struct {
	repo mq.MQProducerRepository
}

func (s *KafkaProducerService) Push(ctx context.Context, msg []byte) error {
	return s.repo.Push(ctx, msg)
}

func (s *KafkaProducerService) BatchPush(ctx context.Context, msgs [][]byte) error {
	return s.repo.BatchPush(ctx, msgs)
}

var _ mq.MQProducerService = (*KafkaProducerService)(nil)

func NewKafkaProducerService(repo mq.MQProducerRepository) *KafkaProducerService {
	return &KafkaProducerService{repo: repo}
}
