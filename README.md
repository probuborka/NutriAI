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
├── Makefile                                   # Makefile
├── prometheus.yml                             # prometheus config
├── promtail-config.yaml                       # promtail config
└── README.md                                  # Документация проекта
```

## Функционал

## Требования

Убедитесь, что следующие инструменты установлены:

- [Go 1.22+](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Сборка

<details>
  <summary>Настройте Dockerfile файл</summary>

```bash  
ENV NUTRIAI_PORT=8080

ENV API_KEY=<your_key_gigachat>

ENV REDIS_HOST=redis

ENV REDIS_PORT=6379

ENV LOG_FILE=./var/log/app.log
```
 </details>

1. Клонируйте репозиторий:

    ```bash
    git clone git@github.com:probuborka/NutriAI.git
    ```

2. Соберите Docker-образ приложения:

    ```bash
    make build
    ```

3. Запустите сервисы с помощью Docker Compose:

    ```bash
    make run-local
    ```
## Команды Makefile

<details>
  <summary>Открыть список команд Make</summary>

- **Собрать Docker-образ приложения**:

    ```bash
    make build
    ```

- **Запустить все сервисы с использованием docker-compose**:

    ```bash
    make run-local
    ```

- **Остановить и удалить все контейнеры**:

    ```bash
    make down
    ```

- **Перезапустить все контейнеры**:

    ```bash
    make restart
    ```

</details>


## [Get recommendation](http://localhost:8080/api/recommendation)

<details>
  <summary>Пример тела запроса (json)</summary>

```json
{
  "user_id": "123456789",
  "user_name": "Евгений",
  "user_data": {
    "profile": {
      "age": 39,
      "gender": "male", // варианты: female male
      "weight_kg": 140,
      "height_cm": 186,
      "fitness_level": "beginner" // варианты: beginner intermediate advanced
    },
    "goals": {
      "primary_goal": "weight_loss", // варианты: weight_loss muscle_toning maintenance
      "secondary_goal": "muscle_toning", // варианты: weight_loss muscle_toning maintenance
      "target_weight_kg": 90,
      "timeframe_weeks": 40
    },
    "preferences": {
      "diet_type": "balanced", // варианты: vegan keto low_carb balanced
      "allergies": ["орехи", "моллюски"], // варианты: перечисление
      "preferred_cuisines": ["средиземноморский", "азиатский"], // варианты: перечисление
      "workout_preferences": ["йога", "силовая тренировка", "кардио"]  // варианты: перечисление
    },
    "lifestyle": {
      "activity_level": "moderate", // варианты: sedentary, light, moderate, active, very_active
      "daily_calorie_intake": 1800,
      "workout_availability_days_per_week": 4,
      "average_sleep_hours": 7
    },
    "medical_restrictions": {
      "has_injuries": true,
      "injury_details": ["травма колена"], // варианты: перечисление
      "chronic_conditions": ["none"]
    }
  },
  "request_details": {
    "service_type": "fitness_nutrition_recommendations",
    "output_format": "weekly_plan", // варианты: daily_plan, weekly_plan, general_advice
    "language": "ru" // варианты: ru, en
  }
}
```

</details>