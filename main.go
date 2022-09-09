package main

import (
	"github.com/icechen128/data-center/internal/app/api_server"
	"github.com/icechen128/data-center/logger"
)

func main() {
	logger.Logger.Info("start server...")
	panic(api_server.RunServer())
}
