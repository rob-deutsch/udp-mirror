package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const version string = "0.1-rob-deutsch"

var (
	listenAddress = flag.String("listen-address", ":9999", "UDP port to listen for incoming packets")
	bufferSize    = flag.Int("buffer-size", 1024, "size of read buffer")
	debug         = flag.Bool("debug", false, "debug mode")
	showVersion   = flag.Bool("version", false, "show version info")
)

type receiver struct {
	address string
	channel chan []byte
}

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: udp-mirror [ ... ]\n\nParameters:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	startServer()
}

func printVersion() {
	fmt.Println("udp-mirror")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author: Rob Deutsch (original author Daniel Czerwonk)")
	fmt.Println("Source code: https://github.com/rob-deutsch/udp-mirror")
}

func startServer() {
	log.Println("Starting listening")
	conn, err := net.ListenPacket("udp", *listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("Listening on %s. Waiting for packets", *listenAddress)

	for {
		buf := make([]byte, *bufferSize)
		len, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println(err)
		}

		if *debug {
			log.Printf("Received (%d)", len)
		}

		conn.WriteTo(buf[:len], addr)
	}
}
