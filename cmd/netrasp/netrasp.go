package main

import (
	"context"

	"github.com/bersennaidoo/netops/application/rest/handlers"
	"github.com/bersennaidoo/netops/application/rest/mid"
	"github.com/bersennaidoo/netops/application/rest/server"
	"github.com/bersennaidoo/netops/infrastructure/device"
	"github.com/bersennaidoo/netops/physical/cisco"
	"github.com/bersennaidoo/netops/physical/config"
	"github.com/bersennaidoo/netops/physical/logger"
)

func main() {
	log := logger.New()
	filename := config.GetConfigFileName()
	cfg := config.New(filename)

	dev := cisco.CreateConfig("R2", "bersen", "bersen")
	defer dev.Close(context.Background())

	dvc := device.New(dev)
	h := handlers.New(log, dvc)
	m := mid.New(log)

	srv := server.New(h, cfg, log, m)
	srv.InitRouter()

	log.Info("Starting the application...")
	srv.Start()
}
