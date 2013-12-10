package network

import (
	"log"
)

type Udp struct{}

func (u Udp) Receive() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":")
	if err != nil {
		log.Fatalf("ResolveUDPAddr failed: %s\n", err)
	}
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("ListenUDP failed: %s\n", err.Error())
	}

	defer listener.Close()

	for {
		message := make([]byte, 1600)
		n, _, err := listener.ReadFromUDP(message)
		log.Printf("Got %d bytes\n%s\n", n, message)

		if err != nil || n == 0 {
			log.Fatalf("Error is: %s, bytes are: %d", err, n)
			continue
		}
	}
}
