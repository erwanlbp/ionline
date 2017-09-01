package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/config"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
)

func main() {
	var err error

	fmt.Println("Start server", config.ServerHost())

	// Init Firebase
	err = sysauth.InitFirebase()
	if err != nil {
		fmt.Println("ERROR Initializing Firebase:", err.Error())
		os.Exit(1)
	}

	// Init Google Auth
	err = extdep.GoogleAuthClient.InitGoogleAuthCredentials()
	if err != nil {
		fmt.Println("ERROR Initializing Google Auth Client:", err.Error())
		os.Exit(1)
	}

	router := handler.Init()
	err = http.ListenAndServe(":"+config.ServerPort(), router)
	if err != nil {
		fmt.Println("Server crashed :", err.Error())
	}
}
