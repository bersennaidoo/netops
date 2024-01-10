package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/netops/domain/models"
	"github.com/bersennaidoo/netops/infrastructure/device"
	"github.com/bersennaidoo/netops/physical/cisco"
	"github.com/kataras/golog"
)

type Handler struct {
	log *golog.Logger
}

func New(log *golog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var configuration models.Configuration

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&configuration)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	dev := cisco.CreateConfig(configuration.Hostname,
		configuration.Username, configuration.Password)
	defer dev.Close(context.Background())

	dvc := device.New(dev)

	output := dvc.CreateConfig(configuration.Config)

	fmt.Fprint(w, output)
}
