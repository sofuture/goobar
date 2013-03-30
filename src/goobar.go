package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"server"
)

const Version = `0.0.1`

var (
	listenAddress   = flag.String("l", "127.0.0.1:22122", "the address to bind to")
	memcacheBackend = flag.String("m", "127.0.0.1:11211", "the memcached backend")
	showVersion     = flag.Bool("v", false, "display version")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *showVersion {
		fmt.Println("goobar", Version)
		os.Exit(0)
	}

	if *listenAddress == "" {
		fmt.Fprintln(os.Stderr, "you must supply an address to listen on")
		flag.Usage()
		os.Exit(1)
	}

	if *memcacheBackend == "" {
		fmt.Fprintln(os.Stderr, "you must specify a memcached backend")
		flag.Usage()
		os.Exit(1)
	}

	log.SetPrefix("goobar ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	tsock, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		panic(err)
	}

	server.Main(tsock)
}
