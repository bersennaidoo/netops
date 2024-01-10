package cisco

import (
	"context"
	"log"

	"github.com/networklore/netrasp/pkg/netrasp"
)

func CreateConfig(hostname string, username string, password string) netrasp.Platform {
	device, err := netrasp.New(hostname,
		netrasp.WithUsernamePassword(username, password),
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
		log.Printf("unable to connect: %v", err)
	}

	return device
}
