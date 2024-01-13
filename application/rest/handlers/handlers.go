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
	"github.com/networklore/netrasp/pkg/netrasp"
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
	var output []netrasp.ConfigResult

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&configuration)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}

	for _, hname := range configuration.Hostname {
	inner:
		for _, uname := range configuration.Username {
			for _, pwd := range configuration.Password {
				dev := cisco.CreateConfig(hname, uname, pwd)
				dvc := device.New(dev)
				for k, config := range configuration.Configs {
					delete(configuration.Configs, k)
					output = append(output, dvc.CreateConfig(config))

					dev.Close(context.Background())
					break inner
				}
			}
		}
	}

	fmt.Fprint(w, output)
}
