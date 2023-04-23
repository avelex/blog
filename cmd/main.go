package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/avelex/blog/internal/app"
	"github.com/avelex/blog/internal/config"
	"github.com/avelex/blog/internal/logger"
)

var (
	insecure bool
)

func init() {
	flag.BoolVar(&insecure, "k", false, "-k")
	flag.Parse()
}

func main() {
	logger := logger.GetLogger()

	config := config.GetConfig()
	config.InsecureHTTP = insecure

	logger.Debugf("app host: %v", config.Host)
	logger.Debugf("app server port: %v", config.HttpPort)
	logger.Debugf("insecure mode: %v", config.InsecureHTTP)
	logger.Debugf("debug mode: %v", config.Debug)

	ctx, cancel := signal.NotifyContext(context.Background(),
		os.Interrupt,
	)

	defer cancel()

	logger.Info("starting app")

	if err := app.NewApp(logger, config).Run(ctx); err != nil {
		logger.Error(err)
	}
}
