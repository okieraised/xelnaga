package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/okieraised/xelnaga/internal/constant"
	"github.com/okieraised/xelnaga/pkg/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	logger *zap.Logger
	engine *gin.Engine
)

func main() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	apiServer := &http.Server{
		Addr:    "0.0.0.0:11104",
		Handler: engine,
	}

	go func() {
		err := apiServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	<-exitSignal
	logger.Info("Received termination signal. Attempting to stop server in 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), constant.TerminationTimeout)
	defer cancel()

	err := apiServer.Shutdown(ctx)
	if err != nil {
		logger.Error(err.Error())
		apiServer.Close()
		return
	}

	select {
	case <-ctx.Done():
		logger.Info("server has been closed")
	}

	logger.Info("Shutdown now. Bye!!!")
}

func init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	engine = gin.New()
	engine.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "Accept", "X-Requested-With", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api.NewInstancesRoute(engine, logger)
}
