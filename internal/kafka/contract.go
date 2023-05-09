package kafka

import "context"

type KafkaMsgHandler func(ctx context.Context, msg []byte) error
