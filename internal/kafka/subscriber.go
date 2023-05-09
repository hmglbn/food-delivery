package kafka

import (
	"context"
	"fmt"
	"food-delivery/internal/helper"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func SubscribeWithHandler(ctx context.Context, r *kafka.Reader, handler KafkaMsgHandler) {
	go func() {
		for {
			if ctx.Err() == context.Canceled {
				// application is closing
				break
			}

			m, err := r.FetchMessage(ctx)
			if err != nil {
				log.Fatalln("FetchMessage problem:", err)
			}

			msgCtx := context.WithValue(
				context.Background(),
				ctxMsgKey,
				fmt.Sprintf("%v/%v/%v-%s", m.Topic, m.Partition, m.Offset, string(m.Key)),
			)
			handlerErr := handler(msgCtx, m.Value)
			if handlerErr != nil {
				helper.Log(fmt.Sprintf("Warning! Handler problem: %s", handlerErr.Error()))
			} else {
				if err := r.CommitMessages(ctx, m); err != nil {
					log.Fatal("failed to commit messages:", err)
				}
			}

			time.Sleep(time.Millisecond * 10)
		}
	}()
}
