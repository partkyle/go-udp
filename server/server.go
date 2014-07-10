package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	host = flag.String("host", "127.0.0.1", "host to listen on")
	port = flag.Int("port", 0, "port to listen on")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*host)
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: *port})
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("listening on %s", listener.LocalAddr())

	io.Copy(os.Stdout, listener)
}
