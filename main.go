package main

import (
	"fmt"
	"net/http"

	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/config"
)

func main() {
	fmt.Println("Start server", config.ServerHost())

	router := handler.Init()
	err := http.ListenAndServe(":"+config.ServerPort(), router)
	if err != nil {
		fmt.Println("Server crashed :", err.Error())
	}
}
