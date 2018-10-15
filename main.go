package main

import (
	"fmt"
	"log"
	"os"

	nats "github.com/nats-io/go-nats"
	"github.com/voutasaurus/env"
)

var (
	service = "synth"
	version = "dev"
)

func main() {
	logger := log.New(os.Stderr, fmt.Sprintf("%s[%s]: ", service, version), log.LstdFlags|log.LUTC|log.Llongfile)
	logger.Println("starting...")

	fatal := func(key string) {
		logger.Fatalf("%q must be set", key)
	}

	var (
		uri = env.Get("SYNTH_NATS_URI").Required(fatal)
	)

	conn, err := nats.Connect(uri)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Printf("server connected to: %q", uri)

	s := &server{
		log:     logger,
		conn:    conn,
		service: service,
	}

	s.subscribe("subject", s.serveSubject)

	// block forever. TODO: implement graceful shutdown
	select {}
}

type server struct {
	log     *log.Logger
	conn    *nats.Conn
	service string
}

func (s *server) subscribe(subject string, fn func(m *nats.Msg)) {
	_, err := s.conn.QueueSubscribe(subject, s.service, fn)
	if err != nil {
		// Note that this only throws an error if the connection is
		// nil, closed, or draining, or if the fn is nil.
		s.log.Fatalf("cannot subscribe: %v", err)
	}
	s.log.Printf("listening to %q as %q", subject, s.service)
}

func (s *server) serveSubject(m *nats.Msg) {
}
