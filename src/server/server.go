package server

import (
	"fmt"
	"log"
	"net"
	"syscall"
)

func Main(listener net.Listener) {
	// var w bool
	for {
		c, err := listener.Accept()
		if err != nil {
			if err == syscall.EINVAL {
				break
			}
			if e, ok := err.(*net.OpError); ok && e.Err == syscall.EINVAL {
				break
			}
			log.Println(err)
			continue
		}

		go serve(c)
	}
}

func serve(nc net.Conn) {
	fmt.Println(nc.RemoteAddr().String())
}
