package main

import (
	"net"
	"log"
	"time"
)

var payload = []byte{194, 91, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 6, 103, 111, 111, 103, 108, 101,
					 3, 99, 111, 109, 0, 0, 1, 0, 1, 0, 0, 41, 16, 0, 0, 0, 0, 0, 0, 0}

func doRequest(throttler chan int){
	defer func(){
		<- throttler
	}()

	conn, err := net.Dial("udp", "0.0.0.0:1054")
	if err != nil {
		log.Println(err)
		return
	}

	count, err := conn.Write(payload)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Send ", count)

	buffer := make([]byte, 1024)

	count, err = conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Read %d", count)
}

func main() {
	throttler := make(chan int, 50)
	count_requests := 1000

	start := time.Now()
	for i:= 0; i < count_requests; i++ {
		throttler <- 1
		go doRequest(throttler)
	}

	d := time.Now().Sub(start).Seconds()

	// 2018/11/12 00:45:46 Send 1000 requests in 16.042471801 seconds. Speed = 62.334534.02 OPS
	log.Printf("Send %d requests in %v seconds. Speed = %f.02 OPS",
		count_requests, d, float64(count_requests)/d)

}
