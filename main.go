package main

import (
	"github.com/adenapila/market-go/app"
	"github.com/adenapila/market-go/logger"
)

func main() {

	//log.Println("starting our application...")
	logger.Info("Starting the application...")
	app.Start()

}
