package main

import (
	desc "github.com/maximpontryagin/auth/pkg/auth_v1"
)

const grpcPort = 50060

type server struct {
	desc.UnimplementedAuthV1Server
}
