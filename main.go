package main

import (
	"fmt"

	"github.com/clo3olb/chaostheory_task1/server"
)

func main() {
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
