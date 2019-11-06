package main

import (
	"fmt"

	"github.com/zhaolion/bus-demo/pkg/biz3"
	_ "github.com/zhaolion/bus-demo/pkg/biz4"
	"github.com/zhaolion/bus-demo/pkg/biz5"
	"github.com/zhaolion/bus-demo/pkg/biz6"
	_ "github.com/zhaolion/bus-demo/pkg/kafka"
)

func main() {
	dispatchMode()
	publishMode()
	fuzzMode()
}

func dispatchMode() {
	fmt.Println("=========== dispatch ===========")
	biz3.Biz("hello")
	fmt.Println("=========== dispatch ===========")
}

func publishMode() {
	fmt.Println("=========== publish ===========")
	biz5.Biz("hello")
	fmt.Println("=========== publish ===========")
}

func fuzzMode() {
	fmt.Println("=========== fuzz ===========")
	biz6.Biz("hello")
	fmt.Println("=========== fuzz ===========")
}
