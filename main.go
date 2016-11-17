package main

import (
	"fmt"
	goredis "gopkg.in/redis.v4"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var out []byte

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
		msgi, err := pubsub.ReceiveTimeout(time.Second * 1000000000)
		if err != nil {
			log.Fatal("receiveMessage", err)
		}

		msg, err := unwrapMessage(msgi)
		if err != nil {
			log.Println(err)
		} else {
			if msg != nil {
				cmd := command
				args := []string{msg.Channel, msg.Payload}
				if out, err = exec.Command(cmd, args...).Output(); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				log.Printf("%s %s", msg.Channel, msg.Payload)
				log.Printf("%s", out)
			}
		}
	}
}

func unwrapMessage(msgi interface{}) (*goredis.Message, error) {

	switch msg := msgi.(type) {
	case *goredis.Subscription:
		// ignore
	case *goredis.Pong:
		// ignore
	case *goredis.Message:
		return msg, nil
	default:
		return nil, fmt.Errorf("no message")

	}
	return nil, nil
}
