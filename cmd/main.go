package main

import (
	"net"
	"log"
	"encoding/hex"
)

func main() {
	listener, err := net.ListenPacket("udp", "0.0.0.0:1054")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1)
	count, _, err := listener.ReadFrom(buffer)
	if err != nil{
		log.Fatal(err)
	}

	log.Println(count)
	log.Println(hex.Dump(buffer))
}
