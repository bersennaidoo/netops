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

	output, err := device.Run(context.Background(), "show version")
	if err != nil {
		log.Fatalf("unable to run command: %v", err)
	}
	fmt.Println(output)
}
