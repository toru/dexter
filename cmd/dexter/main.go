package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting dexter")
	for range time.Tick(time.Second) {
		fmt.Printf("tick: %d\n", time.Now().Unix())
	}
}
