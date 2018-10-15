package main

import (
	"log"
	"os"
	"time"

	nats "github.com/nats-io/go-nats"
	"github.com/voutasaurus/env"
)

func main() {
	logger := log.New(os.Stderr, "pub: ", log.LstdFlags|log.Llongfile|log.LUTC)
	fatal := func(key string) {
		logger.Fatal(key)
	}

	var (
		uri = env.Get("SYNTH_NATS_URI").Required(fatal)
	)

	conn, err := nats.Connect(uri)
	if err != nil {
		logger.Fatal(err)
	}

	msg := []byte("Hello NATS!")

	for {
		if err := conn.Publish("subject", []byte(msg)); err != nil {
			logger.Fatal(err)
		}
		logger.Println("Published")
		time.Sleep(1 * time.Second)
	}
}
