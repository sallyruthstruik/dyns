package main

import (
	"net"
	"log"
)

func askGoogle(request []byte) (response []byte, err error){
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}

	_, err = conn.Write(request)
	if err != nil {
		return
	}

	buffer := make([]byte, 1024)
	count, err := conn.Read(buffer)
	if err != nil {
		return
	}
	return buffer[:count], nil
}

func main() {
	listener, err := net.ListenPacket("udp", "0.0.0.0:1054")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()
	buffer := make([]byte, 1024)

	for {
		count, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		buffer = buffer[:count]
		log.Printf("New NS request: %v", buffer)

		g, err := askGoogle(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response: %v", g)
		_, err = listener.WriteTo(g, addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
