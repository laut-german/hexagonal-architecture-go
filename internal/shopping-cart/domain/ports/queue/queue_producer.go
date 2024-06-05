
package queue

import (
    "context"
)

type QueueProducer interface {
    Publish(ctx context.Context, message []byte) error
}
