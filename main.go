package main

import (
	"log"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 4 {
		usage()
	}

	host := os.Args[1]
	keys := os.Args[2]
	command := os.Args[3]

	r, err := parseURI(host)
	if err != nil {
		log.Fatal("uri\n", err)
	}

	log.Printf("%v %v %v %v", host, keys, command, r)
	pubsub, err := r.client.Subscribe(keys)
	if err != nil {
		log.Fatal("subscribe", err)
	}
	for {
		msg, err := pubsub.ReceiveTimeout(time.Second * 1000000000)
		if err != nil {
			log.Fatal("receiveMessage", err)
		}
		log.Printf("%v", msg)
	}
}
