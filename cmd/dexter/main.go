package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "Path to the config file (required)")
	flag.Parse()

	if len(cfgPath) == 0 {
		flag.PrintDefaults()
		log.Fatal()
	}

	fmt.Println("starting dexter")
	for range time.Tick(time.Second) {
		fmt.Printf("tick: %d\n", time.Now().Unix())
	}
}
