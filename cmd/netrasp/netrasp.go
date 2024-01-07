package main

import (
	"context"
	"fmt"
	"log"

	"github.com/networklore/netrasp/pkg/netrasp"
)

func main() {
	device, err := netrasp.New("192.168.122.195",
		netrasp.WithUsernamePassword("bersen", "bersen"),
		netrasp.WithDriver("ios"),
		netrasp.WithInsecureIgnoreHostKey(),
		netrasp.WithSSHCipher("aes128-cbc"),
		netrasp.WithSSHPort(22),
		netrasp.WithSSHKeyExchange("diffie-hellman-group1-sha1"),
	)

	if err != nil {
		log.Fatalf("unable to initialize device: %v", err)
	}

	err = device.Dial(context.Background())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer device.Close(context.Background())

	config0 := []string{
		"enable",
		"bersen",
		"conf t",
		"int loop 0",
		"ip address 1.1.1.1 255.255.255.255",
		"end",
	}
	output0, err := device.Configure(context.Background(), config0)
	if err != nil {
		log.Fatalf("unable to configure device: %v", err)
	}
	fmt.Println(output0)

	config1 := []string{
		"enable",
		"bersen",
		"conf t",
		"int loop 1",
		"ip address 2.2.2.2 255.255.255.255",
		"end",
	}
	output1, err := device.Configure(context.Background(), config1)
	if err != nil {
		log.Fatalf("unable to configure device: %v", err)
	}
	fmt.Println(output1)

}
