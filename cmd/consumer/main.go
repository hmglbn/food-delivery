package main

import (
	"context"
	"errors"
	"food-delivery/cmd"
	"food-delivery/internal/helper"
	"food-delivery/internal/kafka"
	"food-delivery/internal/refrigerator"
	"log"
	"math/rand"
	"time"
)

func main() {
	mainCtx, mainCtxCancel := context.WithCancel(context.Background())
	rand.Seed(time.Now().UnixNano())

	r := kafka.NewReader(
		mainCtx,
		[]string{cmd.KafkaAddr},
		cmd.KafkaLogin,
		cmd.KafkaPassword,
		cmd.KafkaGroup,
		cmd.KafkaTopic,
	)
	if r == nil {
		log.Fatalln("kafka reader is nil", errors.New("nil reader"))
	}
	kafka.SubscribeWithHandler(mainCtx, r, refrigerator.GotNewFoodHandler)

	helper.ListenSigInt()
	helper.CloseApp(mainCtxCancel)
}
