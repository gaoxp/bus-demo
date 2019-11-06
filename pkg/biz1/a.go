package biz1

import (
	"fmt"

	"github.com/zhaolion/demo/bus"
	"github.com/zhaolion/demo/model"
)

func Biz(msg string) {
	err := bus.Dispatch(&model.Message{Content: msg})
	fmt.Printf("[biz1] dispatch msg %v with err: %v\n", msg, err)
}
