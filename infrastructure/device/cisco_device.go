package device

import (
	"context"
	"log"

	"github.com/networklore/netrasp/pkg/netrasp"
)

type Device struct {
	device netrasp.Platform
}

func New(device netrasp.Platform) *Device {
	return &Device{
		device: device,
	}
}

func (d *Device) CreateConfig(config []string) netrasp.ConfigResult {

	output, err := d.device.Configure(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to configure device: %v", err)
	}

	return output
}
