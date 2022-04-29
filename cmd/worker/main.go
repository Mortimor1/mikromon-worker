package main

import (
	"github.com/Mortimor1/mikromon-worker/internal/config"
	"github.com/Mortimor1/mikromon-worker/internal/webserver"
	"github.com/Mortimor1/mikromon-worker/pkg/logging"
)

func main()  {

	logger := logging.GetLogger()

	cfg := config.GetConfig()

	server := new(webserver.Server)

	if err := server.Run(cfg); err != nil {
		logger.Fatalf("error running http server: %s", err.Error())
	}
}
