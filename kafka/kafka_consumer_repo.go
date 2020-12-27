package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/t/mq"
)

type KafkaConsumerRepository struct {
	r *kafka.Reader
}

func (s *KafkaConsumerRepository) Fetch(ctx context.Context) ([]byte, error) {
	msg, err := s.r.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Partition: ", msg.Partition, ", Offset: ", msg.Offset, "Value: ", string(msg.Value))
	return msg.Value, nil
}

var _ mq.MQConsumerRepository = (*KafkaConsumerRepository)(nil)

func NewKafkaConsumerRepository(brokers []string, group, topic string) *KafkaConsumerRepository {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: group,
		Topic:   topic,
	})
	return &KafkaConsumerRepository{r: r}
}
