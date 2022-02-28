package main

import (
	pb "GHZExample/internal/proto_db"
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = 50052
)

type server struct {
	pb.UnimplementedGreeterServer
}

func debug(f string, s string) {
	log.Printf("[%s]: %v", f, s)
}

func (s *server) SayHello(ctx context.Context, in *pb.ServerRequest) (*pb.ServerResponse, error) {
	f := "SayHello"
	debug(f, in.Name)
	res := fmt.Sprintf("Hello %s!", in.GetName())
	return &pb.ServerResponse{Res: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer func() { //cleanup
		s.Stop()
		if err := lis.Close(); err != nil {
			log.Println(err)
		}
	}()
}
