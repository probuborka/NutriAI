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

	// Настройка Logrus
	log := logrus.New()

	// Настройка формата вывода (JSON)
	log.SetFormatter(&logrus.JSONFormatter{})

	// Настройка вывода в файл
	file, err := os.OpenFile("./var/log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("log file error")
		return
	}
	log.SetOutput(file)

	//
	cfg, err := config.New()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("error config")
		return
	}

	//gigachat client
	gigaChatClient := gigachatclient.New(
		cfg.Api.Key,
	)

	//redis client
	redisClient := redisclient.NewClient(
		&redisclient.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

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

	//service
	useCaseMetric := metric.NewMetricUseCase(
		prometheus,
	)

	useCaseRecommendation := recommendation.NewRecommendationUseCase(
		gigaChatRecommendation,
		cacheRecommendation,
	)

	//handlers
	handlers := handlers.New(
		useCaseRecommendation,
		useCaseMetric,
		log,
	)

	//http server
	server := route.New(
		cfg.HTTP.Port,
		handlers.Init(),
	)

	//start server
	//logger.Info("server start, port:", cfg.HTTP.Port)
	log.WithFields(logrus.Fields{
		"service": "nutrial",
		"version": "1.0.0",
		"port":    cfg.HTTP.Port,
	}).Info("Server run")
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			//logger.Errorf("error occurred while running http server: %s\n", err.Error())
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
		//logger.Errorf("failed to stop server: %v", err)
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to stop server")
	}
}
