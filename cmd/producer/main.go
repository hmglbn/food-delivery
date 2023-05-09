package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"food-delivery/cmd"
	"food-delivery/internal/helper"
	"food-delivery/internal/kafka"
	"food-delivery/internal/refrigerator"
	"log"
	"math/rand"
	"time"
)

func main() {
	_, mainCtxCancel := context.WithCancel(context.Background())
	rand.Seed(time.Now().UnixNano())

	w := kafka.NewWriter(
		[]string{cmd.KafkaAddr},
		cmd.KafkaLogin,
		cmd.KafkaPassword,
		cmd.KafkaTopic,
	)
	if w == nil {
		log.Fatalln("kafka writer is nil", errors.New("nil writer"))
	}

	sendFood(w)

	helper.CloseApp(mainCtxCancel)
}

func sendFood(w *kafka.Writer) {
	food := refrigerator.Food{
		Name:   &refrigerator.Names[rand.Intn(len(refrigerator.Names))],
		Volume: rand.Intn(25),
		Weight: rand.Intn(50),
	}
	helper.Log(fmt.Sprintf("sending: %s (%dL, %dkg)", *food.Name, food.Volume, food.Weight))
	foodJson, _ := json.Marshal(food)
	if err := w.Write(foodJson); err != nil {
		log.Fatalln("problem with writing", err)
	}
}
