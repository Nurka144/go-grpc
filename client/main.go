package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/go-grpc/pkg/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5051"
	x       = 1
	y       = 2
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewAdderClient(conn)

	default_x := x
	default_y := y

	if len(os.Args) > 2 {
		default_x, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("err in x")
		}
		default_y, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("err in y")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{X: int32(default_x), Y: int32(default_y)})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add: %v", r.GetResult())
}
