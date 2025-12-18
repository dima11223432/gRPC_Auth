package app

import (
	grpcApp "authService/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCapp *grpcApp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcapp := grpcApp.New(log, grpcPort)
	return &App{
		GRPCapp: grpcapp,
	}

}
