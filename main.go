package main

import (
	"fmt"

	"github.com/ErwanLBP/IOnline/internal"
	"github.com/ErwanLBP/IOnline/internal/sys/sysconf"
)

func main() {
	fmt.Println("Hello world !")
	fmt.Println("Host:", sysconf.ServerHost())

	internal.InitServer()
}
