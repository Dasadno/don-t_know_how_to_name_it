package main

import (
	"context"

	authv1 "github.com/Dasadno/service/server/gen/auth/v1"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkpostgres "github.com/rookie-ninja/rk-db/postgres"
	rkredis "github.com/rookie-ninja/rk-db/redis"
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
	boot := rkboot.NewBoot(rkboot.WithBootConfigPath("services/auth-service/boot.yaml", nil))

	boot.Bootstrap(context.TODO())

	// REDIS
	// auto migrate database and init userDb variable
	redisEntry := rkredis.GetRedisEntry("server")
	redisClient, _ := redisEntry.GetClient()
	_ = redisClient

	// POSTGRES
	// auto migrate database and init userDb variable
	pgEntry := rkpostgres.GetPostgresEntry("auth-db")
	userDb := pgEntry.GetDB("user")
	if !userDb.DryRun {
		userDb.AutoMigrate(&User{})
	}

	// GRPC
	grpcEntry := rkgrpc.GetGrpcEntry("auth-service")

	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		authv1.RegisterAuthServiceServer(server, &AuthServer{})
	})
	grpcEntry.AddRegFuncGw(authv1.RegisterAuthServiceHandlerFromEndpoint)

	boot.WaitForShutdownSig(context.TODO())
}

type User struct {
	// todo
}
