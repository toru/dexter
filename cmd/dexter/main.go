package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "Path to the config file (required)")
	flag.Parse()

	if len(cfgPath) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("starting dexter")
	for range time.Tick(time.Second) {
		fmt.Printf("tick: %d\n", time.Now().Unix())
	}
}
