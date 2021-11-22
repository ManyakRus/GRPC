package main

import (
	pb "GRPC/.GRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"time"
)

//var server pb.UnimplementedGreeterServer

//type servers struct{
//	server pb.UnimplementedGreeterServer
//}

type servers pb.UnimplementedGreeterServer

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	server1 := pb.UnimplementedGreeterServer{}

	pb.RegisterGreeterServer(grpcServer, &server1)

	println("Server started ", time.Now().String())

	grpcServer.Serve(listener)

}

//func (s *servers) SayHello(c context.Context, request *pb.HelloRequest)(response *pb.HelloReply, err error) {
//n := 0
//// Сreate an array of runes to safely reverse a string.
//rune := make([]rune, len(request.Name))
//
//for _, r := range request.Name {
//rune[n] = r
//n++
//}
//
//// Reverse using runes.
//rune = rune[0:n]
//
//for i := 0; i < n/2; i++ {
//rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
//}
//
//output := string(rune)
//response = &pb.HelloReply{
//Message: output,
//}
//
//return response, nil
//}
