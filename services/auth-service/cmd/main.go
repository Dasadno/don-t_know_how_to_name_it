package main

import (
	"context"

	authv1 "github.com/Dasadno/service/server/gen/auth/v1"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

type AuthServer struct {
	authv1.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.AuthResponse, error) {
	return &authv1.AuthResponse{AccessToken: "fake-jwt"}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.AuthResponse, error) {
	return &authv1.AuthResponse{AccessToken: "fake-jwt"}, nil
}

func main() {
	ctx := context.WithoutCancel(context.Background())
	boot := rkboot.NewBoot(rkboot.WithBootConfigPath("services/auth-service/boot.yaml", nil))
	entry := rkgrpc.GetGrpcEntry("auth-service")

	entry.AddRegFuncGrpc(func(server *grpc.Server) {
		authv1.RegisterAuthServiceServer(server, &AuthServer{})
	})
	entry.AddRegFuncGw(authv1.RegisterAuthServiceHandlerFromEndpoint)

	boot.Bootstrap(ctx)
	boot.WaitForShutdownSig(ctx)
}
