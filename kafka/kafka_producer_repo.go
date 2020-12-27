package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/t/mq"
)

type KafkaProducerRepository struct {
	w      *kafka.Writer
	broker string
	topic  string
}

func (s *KafkaProducerRepository) Push(ctx context.Context, b []byte) error {
	msg := kafka.Message{
		Value: b,
	}
	return s.w.WriteMessages(ctx, msg)
}

func (s *KafkaProducerRepository) BatchPush(ctx context.Context, bs [][]byte) error {
	var msgs []kafka.Message

	for _, b := range bs {
		m := kafka.Message{
			Value: b,
		}
		msgs = append(msgs, m)
	}
	return s.w.WriteMessages(ctx, msgs...)
}

var _ mq.MQProducerRepository = (*KafkaProducerRepository)(nil)

func NewKafkaProducerRepository(broker, topic string) *KafkaProducerRepository {
	w := &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: topic,
		//RequiredAcks: kafka.RequireAll,
		RequiredAcks: kafka.RequireNone,
	}

	return &KafkaProducerRepository{
		w:      w,
		broker: broker,
		topic:  topic,
	}
}
