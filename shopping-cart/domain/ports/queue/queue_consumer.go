package queue

import (
    "context"
)

type QueueConsumer interface {
    StartConsuming(ctx context.Context, handler func(message []byte) error) error
    HandleMessage(message []byte) error
}
