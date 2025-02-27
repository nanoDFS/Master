package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nanoDFS/p2p/p2p/encoder"
	"github.com/nanoDFS/p2p/p2p/transport"
)

type Sample struct {
	Name string
}

func main() {
	fmt.Println("Hello,")
	server, err := transport.NewTCPTransport(":9000")
	if err != nil {
		log.Fatal("couldn't")
	}
	server.Listen()

	go func() {
		for {
			var msg Sample
			server.Consume(encoder.GOBDecoder{}, &msg)
			log.Println(msg.Name)
		}
	}()

	client, _ := transport.NewTCPTransport(":9029")
	data := Sample{Name: "Nagaraj"}
	client.Send(":9000", data)
	client.Send(":9000", data)
	client.Send(":9000", data)

	time.Sleep(time.Second * 3)
}
