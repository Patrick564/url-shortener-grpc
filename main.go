package main

import (
	"log"
	"net"

	"github.com/Patrick564/url-shortener-grpc/server"
	pb "github.com/Patrick564/url-shortener-grpc/url_shortener"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("error at listen tcp port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUrlSortenerServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
