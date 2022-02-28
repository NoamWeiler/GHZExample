package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
}

func debug(f string, s string) {
	log.Printf("[%s]: %v", f, s)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := NewGRPCServer()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer func() { //cleanup
		adminIsConnected = false
		s.Stop()
		if err := lis.Close(); err != nil {
			log.Println(err)
		}
	}()
}

//TODO
/*
	1)	add signal interrupt for proper closing ctl+c
	2)	start working on the measures handling (incllude DB)
*/
