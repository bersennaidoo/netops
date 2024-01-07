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

func (d *Device) ConfigureDevice(cmd1 string, pwd string, cmd2 string,
	cmd3 string, cmd4 string, cmd5 string) netrasp.ConfigResult {
	config := []string{
		cmd1,
		pwd,
		cmd2,
		cmd3,
		cmd4,
		cmd5,
	}
	output, err := d.device.Configure(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to configure device: %v", err)
	}

	return output
}
