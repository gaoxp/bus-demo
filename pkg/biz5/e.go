package biz5

import (
	"fmt"

	"github.com/zhaolion/demo/bus"
	"github.com/zhaolion/demo/model"
)

func Biz(msg string) {
	err := bus.Publish(&model.Message{Content: msg})
	fmt.Printf("[biz4] publish msg %v with err: %v\n", msg, err)
}
