package kafka

import (
	"fmt"

	"github.com/zhaolion/demo/bus"
	"github.com/zhaolion/demo/model"
)

func init() {
	bus.AddHandler("kafka", Produce)
	bus.AddEventListener(Produce)
}

func Produce(msg *model.Message) error {
	fmt.Printf("[kafka-msg]: content %v\n", msg.Content)
	return nil
}
