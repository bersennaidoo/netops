package main

import (
	"context"
	"fmt"

	"github.com/bersennaidoo/netops/infrastructure/device"
	"github.com/bersennaidoo/netops/physical/cisco"
)

func main() {
	dev := cisco.CreateConfig("192.168.122.195", "bersen", "bersen")
	defer dev.Close(context.Background())

	dvc := device.New(dev)

	output := dvc.ConfigureDevice("enable", "bersen", "conf t",
		"int loop 2", "ip address 3.3.3.3 255.255.255.255", "end")

	fmt.Println(output)
}
