package main

import (
	pb "GHZExample/internal/proto_db"
	context "context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port = 50052
)

var (
	s *grpc.Server
)

type server struct {
	pb.UnimplementedGreeterServer
}

func debug(f string, s string) {
	log.Printf("[%s]: %v", f, s)
}

func (s *server) SayHello(ctx context.Context, in *pb.ServerRequest) (*pb.ServerResponse, error) {
	f, name := "SayHello", in.GetName()
	debug(f, name)
	res := fmt.Sprintf("Hello %s!", name)
	return &pb.ServerResponse{Res: res}, nil
}

func createGRPC() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("gRPC server: failed to listen.\nerror:%v", err)
	}

	s = grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", ln.Addr())
	return s.Serve(ln)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer close(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(createGRPC)

	select {
	case in := <-interrupt:
		fmt.Println(in)
		break
	case <-ctx.Done():
		break
	}

	//put the server up for 10 minutes, then close it gracefully
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer shutdownCancel()

	if s != nil {
		s.GracefulStop()
	}

	err := g.Wait()
	if err != nil {
		log.Fatalf("server returning an error.\nerror:%v", err)
	}

	fmt.Println("exit..")
}
