package mq

import "context"

type MQProducerService interface {
	Push(context.Context, []byte) error
	BatchPush(context.Context, [][]byte) error
}

type MQProducerRepository interface {
	Push(context.Context, []byte) error
	BatchPush(context.Context, [][]byte) error
}
