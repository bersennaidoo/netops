package services

import "github.com/bersennaidoo/netops/domain/models"

type DeviceConfiguration struct {
	Cmds []models.CMD
}

func (dc DeviceConfiguration) CreateConfiguration() []string {

	var config []string

	for _, cmd := range dc.Cmds {
		config = append(config, cmd.Cmd)
	}

	return config
}
