package biz3

import (
	"fmt"

	"github.com/zhaolion/bus-demo/pkg/biz1"
	"github.com/zhaolion/bus-demo/pkg/biz2"
)

func Biz(msg string) {
	fmt.Println("[biz3] call it")
	biz2.Biz(msg)
	biz1.Biz(msg)
}
