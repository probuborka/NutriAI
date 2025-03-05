# NutriAI
## Описание проекта
Микросервис, разработанный на языке Go, который предоставляет пользователям персонализированные рекомендации по питанию и тренировкам

## Структура проекта
```plaintext
.
├── cmd/
│   └── main.go                                # Точка входа в приложение
├── docs/                                      # Swagger
├── integration/                               # Интеграционные тесты
│   │── gigachat_test.go                       # GigaChat тест
│   │── redis_test.go                          # Redis тест
│   └── settings.go                            # Настройки тестов
├── internal/                                  # Основной код приложения
│   ├── app/                                   # app NutriAI
│   │   └── app.go
│   ├── config/                                # Конфигурация
│   │   └── config.go
│   │── controller/                            # Взаимодействие с внешним миром (точка входа в приложение)
│   │   └── http/                              # HTTP-обработчики
│   │       │── handler.go
│   │       │── logging.go                     # Логирование запросов
│   │       │── middleware.go            
│   │       │── recommendation_handler.go      # Get recommendation
│   │       │── record_metrics.go              # Metrics for prometheus
│   │       └── response.go
│   ├── entity/                                # Бизнес-сущности
│   │   │── config_entity.go 
│   │   │── error_entity.go                    # модели данных Error                
│   │   │── metric_entity.go                       
│   │   │── recommendation_entity_test.go      # Unit-тесты Бизнес-правила для модели данных UserRecommendationRequest
│   │   └── recommendation_entity.go           # модели данных Recommendation
│   ├── infrastructure/                        # Взаимодействие с внешними системами (база данных, кеш и т.д.)
│   │   │── gigachat/                          # gigachat
│   │   │   └── recommendation_gigachat.go    
│   │   │── prometheus/                        # prometheus
│   │   │   └── prometheus.go
│   │   └── redis/                             # redis
│   │       └── recommendation_redis.go
│   └── usecase/                               # бизнес-логика
│       │── metric/                           
│       │   └── metric_usecase.go    
│       └── recommendation/                    # бизнес-логика получения рекомендаций
│           │── recommendation_usecase_test.go # Unit-тесты usecase recommendation
│           └── recommendation_usecase.go      # usecase recommendation
├── pkg/
│   ├── gigachat/                              # gigachat клиент
│   │   │── access_token.go
│   │   │── generate_text.go
│   │   └── gigachat.go
│   └── route/                                 # route клиент
│       └── route.go
├── .dockerignore                              # Игнорируемые файлы для Docker
├── .gitignore                                 # Игнорируемые файлы для Git
├── go.mod                                     # Файл зависимостей Go
├── go.sum                                     # Контрольная сумма зависимостей
├── loki-config.yaml                           # loki config
├── prometheus.yml                             # prometheus config
├── promtail-config.yaml                       # promtail config
└── README.md                                  # Документация проекта
```
