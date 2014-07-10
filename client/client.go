package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host = flag.String("host", "127.0.0.1", "host to send udp to")
	port = flag.Int("port", 0, "port to send udp to")

	payload = flag.String("payload", "", "data to send over udp")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)

	srcAddr := &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: *port}

	log.Printf("sending from %s to %s", srcAddr, dstAddr)

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s", *payload)
}
