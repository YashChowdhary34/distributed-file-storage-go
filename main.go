package main

import (
	"log"

	"github.com/YashChowdhary34/distributed-file-storage-go/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: 	":3000",
		ShakeHands: 	p2p.NOPHandshakeFunc,
		Decoder: 			p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
}