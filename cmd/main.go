package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerofglory/vector-study-case/internal/adapter"
	service "github.com/ownerofglory/vector-study-case/internal/core/service"
	"github.com/ownerofglory/vector-study-case/internal/logging"
	"github.com/ownerofglory/vector-study-case/internal/server"
	"go.uber.org/zap"
)

func main() {
	// create and init logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugarLog := logger.Sugar()
	logging.InitLogger(sugarLog)

	// init middleware
	logMiddleware := server.LoggingMiddleware()

	// init services
	numService := service.NewNumService()

	// init gin router
	r := gin.Default()
	r.Use(logMiddleware)
	h := adapter.NewHandler(numService)
	r.GET(server.FindPairForTargetRoute, h.HandleFindPair)

	// start web server
	c := make(chan error)
	go func() {
		c <- r.Run(":80")
		logging.Info("Gracefully shut down the server")
	}()
	logging.Info("Application started")

	<-c
}
