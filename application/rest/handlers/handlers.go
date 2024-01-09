package handlers

import (
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
	var cmds []models.CMD

	for i := 0; i < 1; i++ {
		cmds = append(cmds, models.CMD{"enable"})
		cmds = append(cmds, models.CMD{"bersen"})
		cmds = append(cmds, models.CMD{"conf t"})
		cmds = append(cmds, models.CMD{"int loop 3"})
		cmds = append(cmds, models.CMD{"ip address 4.4.4.4 255.255.255.0"})
		cmds = append(cmds, models.CMD{"end"})
		cmds = append(cmds, models.CMD{"wr"})
	}

	devconfig := services.DeviceConfiguration{
		Cmds: cmds,
	}

	config := devconfig.CreateConfiguration()
	output := h.dev.CreateConfig(config)

	fmt.Fprint(w, output)
}
