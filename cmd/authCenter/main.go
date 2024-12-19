package main

import (
	"flag"
	"fmt"
	"os"
	"smj-go/pkg/autoConfig"
)

var (
	configDir = flag.String("configs", os.Getenv("SMJ_CONFIG"), "Specify configuration directory.(env:SMJ_CONFIG,default:./config)")
)

func main() {
	flag.Parse()
	c := autoConfig.AutoConfig{}
	autoConfig.StartApp(*configDir, c)
	fmt.Println(autoConfig.C)
}
