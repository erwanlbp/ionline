package main

import (
	"fmt"

	"github.com/erwanlbp/ionline/internal"
	"github.com/erwanlbp/ionline/internal/sys/sysconf"
)

func main() {
	fmt.Println("Hello world !")
	fmt.Println("Host:", sysconf.ServerHost())

	internal.InitServer()
}
