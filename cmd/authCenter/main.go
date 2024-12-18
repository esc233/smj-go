package main

import (
	"flag"
	"fmt"
	"os"
	"smj-go/pkg/config"
)

var (
	configDir = flag.String("configs", os.Getenv("SMJ_CONFIG"), "Specify configuration directory.(env:SMJ_CONFIG,default:./config)")
)

func main() {
	flag.Parse()
	config.LoadConfig(*configDir)
	fmt.Println(config.C)
}
