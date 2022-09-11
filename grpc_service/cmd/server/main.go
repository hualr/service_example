package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "service_example/grpc_service/api/message"
)

const (
	port = ":50051"
)

// 新版本 gRPC 要求必须嵌入 pb.UnimplementedGreeterServer 结构体

type messagePro struct {
	pb.UnimplementedHealthServiceServer
}

func (g *messagePro) CheckHealth(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", in.String())
	return &pb.MessageResponse{ResponseSomething: "Hello " + in.SaySomething}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	// 将服务描述(server)及其具体实现(greeterServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
	pb.RegisterHealthServiceServer(server, &messagePro{})
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
