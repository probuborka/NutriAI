package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/probuborka/NutriAI/internal/config"
	handlers "github.com/probuborka/NutriAI/internal/controller/http"
	"github.com/probuborka/NutriAI/internal/infrastructure/gigachat"
	"github.com/probuborka/NutriAI/internal/infrastructure/prometheus"
	"github.com/probuborka/NutriAI/internal/infrastructure/redis"
	"github.com/probuborka/NutriAI/internal/usecase/metric"
	"github.com/probuborka/NutriAI/internal/usecase/recommendation"
	gigachatclient "github.com/probuborka/NutriAI/pkg/gigachat"
	"github.com/probuborka/NutriAI/pkg/route"
	redisclient "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func Run() {
	//config -----------------------------------------------------------------------------------------------------------
	cfg, err := config.New()
	if err != nil {
		logrus.Error(err)
		return
	}

	//log --------------------------------------------------------------------------------------------------------------
	//logrus
	log := logrus.New()

	//format log json
	log.SetFormatter(&logrus.JSONFormatter{})

	//saving logs to file
	file, err := os.OpenFile(cfg.Log.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("log file error")
		return
	}
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	//client ----------------------------------------------------------------------------------------------------------
	//gigachat client
	gigaChatClient := gigachatclient.New(
		cfg.Api.Key,
	)

	//redis client
	redisClient := redisclient.NewClient(
		&redisclient.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
			Password: "",
			DB:       0,
		},
	)

	//infrastructure ---------------------------------------------------------------------------------------------------
	//prometheus
	prometheus := prometheus.NewPrometheus()

	//gigachat
	gigaChatRecommendation := gigachat.NewRecommendation(
		gigaChatClient,
	)

	//cach
	cacheRecommendation := redis.NewRecommendation(
		redisClient,
	)

	//usecase ------------------------------------------------------------------------------------------------------
	useCaseMetric := metric.NewMetricUseCase(
		prometheus,
	)

	useCaseRecommendation := recommendation.NewRecommendationUseCase(
		gigaChatRecommendation,
		cacheRecommendation,
	)

	//handlers ------------------------------------------------------------------------------------------------------
	handlers := handlers.New(
		useCaseRecommendation,
		useCaseMetric,
		log,
	)

	//server --------------------------------------------------------------------------------------------------------
	//http server
	server := route.New(
		cfg.HTTP.Port,
		handlers.Init(),
	)

	//start server
	log.WithFields(logrus.Fields{
		"service": "nutrial",
		"version": "1.0.0",
		"port":    cfg.HTTP.Port,
	}).Info("Server run")
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("error occurred while running http server")
		}
	}()

	//stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.WithFields(logrus.Fields{
		"service": "nutrial",
		"version": "1.0.0",
		"port":    cfg.HTTP.Port,
	}).Info("server stop")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to stop server")
	}
}
