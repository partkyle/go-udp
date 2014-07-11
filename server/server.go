package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	logger = log.New(os.Stdout, "[udp-server] ", log.LstdFlags)

	host = flag.String("host", "0.0.0.0", "host to listen on")
	port = flag.Int("port", 0, "port to listen on")

	blockSize = flag.Int("size", 1024, "block size to read packets on")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: *port})
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.Printf("listening on addr=%s with block size=%d", listener.LocalAddr(), *blockSize)

	data := make([]byte, *blockSize)
	for {
		n, remoteAddr, err := listener.ReadFrom(data)
		if err != nil {
			logger.Fatalf("error during read: %s", err)
		}

		logger.Printf("<%s> %s", remoteAddr, data[:n])
	}
}
