package main

import (
	"fmt"

	"github.com/zhaolion/demo/pkg/biz6"

	"github.com/zhaolion/demo/pkg/biz3"
	_ "github.com/zhaolion/demo/pkg/biz4"
	"github.com/zhaolion/demo/pkg/biz5"
	_ "github.com/zhaolion/demo/pkg/kafka"
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
