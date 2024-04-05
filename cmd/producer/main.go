package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

type message struct {
	Timestamp string  `json:"timestamp"`
	Value     float32 `json:"value"`
}

func main() {
	nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		ts := time.Now().Format(time.RFC3339)
		val := rand.Float32()
		msg := message{ts, val}
		data, _ := json.Marshal(msg)

		nc.Publish("record", data)
	}
}
