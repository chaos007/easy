package main

import (
	"fmt"

	"github.com/chaos007/easycome/session/grpc"
)

func main() {
	g := grpc.NewSession()
	fmt.Println("-------------g:", g)
}
