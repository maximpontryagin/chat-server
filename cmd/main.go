package main

import (
	desc "github.com/maximpontryagin/chat-server/pkg/chat_server_v1"
)

const grpcPort = 50060

type server struct {
	desc.UnimplementedAuthV1Server
}
