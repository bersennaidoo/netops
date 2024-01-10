package main

import (
	"github.com/bersennaidoo/netops/application/rest/handlers"
	"github.com/bersennaidoo/netops/application/rest/mid"
	"github.com/bersennaidoo/netops/application/rest/server"
	"github.com/bersennaidoo/netops/physical/config"
	"github.com/bersennaidoo/netops/physical/logger"
)

func main() {
	log := logger.New()
	filename := config.GetConfigFileName()
	cfg := config.New(filename)

	h := handlers.New(log)
	m := mid.New(log)

	srv := server.New(h, cfg, log, m)
	srv.InitRouter()

	log.Info("Starting the application...")
	srv.Start()
}
