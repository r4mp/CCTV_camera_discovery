package network

import (
	"bufio"
	"log"
	"net"
	"os"
)

type Udp struct{}

// Send("239.255.255.250:3702")

func (u Udp) Send(message string, addr string) (int, error) {

	laddr, err := net.ResolveUDPAddr("udp", ":0")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	maddr, err := net.ResolveUDPAddr("udp4", addr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c, err := net.ListenUDP("udp4", laddr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer c.Close()

	n, err := c.WriteToUDP([]byte(message), maddr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return n, nil

}
