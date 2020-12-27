package mq

import "context"

type MQConsumerService interface {
	Fetch(context.Context) ([]byte, error)
}

type MQConsumerRepository interface {
	Fetch(context.Context) ([]byte, error)
}
