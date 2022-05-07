package main

import (
	pb "JKZDB/db/proto"
	sl "JKZDB/db/server_lib"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 8080, "The server port")

func main() {
	flag.Parse()
	s := grpc.NewServer()
	server, err := sl.MakeJKZDBServer(*port)
	if err != nil {
		log.Fatalf("Could not start DB server.")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterJKZDBServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
