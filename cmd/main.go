package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/maximpontryagin/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50060

type server struct {
	desc.UnimplementedChat_ServerV1Server
}

func (s *server) Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{Id: 1}, nil
}

func (s *server) Delete(context.Context, *desc.DeleteRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *server) SendMessage(context.Context, *desc.SendRequest) (*empty.Empty, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChat_ServerV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
