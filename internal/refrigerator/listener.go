package refrigerator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"food-delivery/internal/helper"
	"math/rand"
	"time"
)

func GotNewFoodHandler(ctx context.Context, msg []byte) error {
	helper.Log("Wow! We got new nice food for refrigerator. Let's check what is it!")

	var food Food
	if err := json.Unmarshal(msg, &food); err != nil {
		return err
	}
	if food.Name == nil {
		return errors.New("food name is empty")
	}

	putFoodInRefrigerator(ctx, food)

	return nil
}

func putFoodInRefrigerator(ctx context.Context, food Food) {
	msgKey := getMsgKeyFromCtx(ctx)
	helper.Log(fmt.Sprintf("it takes a long time to put %s (%dL, %dkg) in the refrigerator (%s)", *food.Name, food.Volume, food.Weight, msgKey))
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(900)+100))
	helper.Log(fmt.Sprintf("we put %s, success (%s)", *food.Name, msgKey))
}

func getMsgKeyFromCtx(ctx context.Context) string {
	msgKey := "empty"
	if msgKeyRaw := ctx.Value(ctxMsgKey); msgKeyRaw != nil {
		msgKey = msgKeyRaw.(string)
	}
	return msgKey
}
