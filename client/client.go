package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	logger = log.New(os.Stdout, "[udp-server] ", log.LstdFlags)

	host = flag.String("host", "0.0.0.0", "host to send udp to")
	port = flag.Int("port", 0, "port to send udp to")

	payload = flag.String("payload", "", "data to send over udp")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)

	srcAddr := &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: *port}

	logger.Printf("sending from %s to %s", srcAddr, dstAddr)

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	var r io.Reader = strings.NewReader(*payload)
	if *payload == "" {
		logger.Printf("using stdin for payload")
		r = os.Stdin
	}

	n, err := io.Copy(conn, r)
	if err != nil {
		logger.Fatal("error sending data: %s", err)
	}

	logger.Printf("sent %d bytes", n)
}
