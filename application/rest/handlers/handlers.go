package handlers

import (
	"fmt"
	"net/http"

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
	fmt.Fprint(w, "hello device")
}
