package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/go-grpc/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	port = ":5051"
)

type server struct {
	pb.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: x => %v; y => %v", in.GetX(), in.GetY())
	return &pb.AddResponse{Result: in.GetX() + in.GetY()}, nil
}

func main() {

	go func() {
		mux := runtime.NewServeMux()

		pb.RegisterAdderHandlerServer(context.Background(), mux, &server{})

		log.Fatalln(http.ListenAndServe("localhost:5052", mux))
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
