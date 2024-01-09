package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/netops/domain/models"
	"github.com/bersennaidoo/netops/domain/services"
	"github.com/bersennaidoo/netops/infrastructure/device"
	"github.com/kataras/golog"
)

type Handler struct {
	log *golog.Logger
	dev *device.Device
}

func New(log *golog.Logger, dev *device.Device) *Handler {
	return &Handler{
		log: log,
		dev: dev,
	}
}

func (h *Handler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var loopbackcmds models.LoopBackRequest

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&loopbackcmds)
	if err != nil {
		http.Error(w, "Could not decode", http.StatusBadRequest)
		return
	}

	cmd := models.CMD{}
	cmds := cmd.CreateCMDList(loopbackcmds)

	devconfig := services.DeviceConfiguration{
		Cmds: cmds,
	}

	config := devconfig.CreateConfiguration()
	output := h.dev.CreateConfig(config)

	fmt.Fprint(w, output)
}
