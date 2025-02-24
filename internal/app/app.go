package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/probuborka/NutriAI/internal/config"
	handlers "github.com/probuborka/NutriAI/internal/controller/http"
	"github.com/probuborka/NutriAI/pkg/logger"
	"github.com/probuborka/NutriAI/pkg/route"
)

func Run() {

	cfg, err := config.New()
	if err != nil {
		//	logger.Error(err)
		return
	}

	//handlers
	handlers := handlers.New()

	//http server
	server := route.New(cfg.HTTP.Port, handlers.Init())

	//start server
	logger.Info("server start, port:", cfg.HTTP.Port)
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	//stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Info("server stop")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
