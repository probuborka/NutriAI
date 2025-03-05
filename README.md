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

# Требования

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


## Get recommendation

```bash
http://localhost:8080/api/recommendation
 ```

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

<details>
  <summary>Пример ответа (json)</summary>

```json
{
    "recommendations": "Евгений, исходя из предоставленной информации, я могу предложить вам следующий план действий для достижения ваших целей.\n\n### Ваша Цель:\n- Потеря веса (Primary Goal)\n- Укрепление мышц (Secondary Goal)\n- Целевой вес: 90 кг\n- Срок реализации: 40 недель\n\n### Индивидуальные Предпочтения и Ограничения:\n- Тип диеты: Сбалансированная диета\n- Аллергии: Орехи, Моллюски\n- Предпочитаемые кухни: Средиземноморская, Азиатская\n- Физическая активность: Йога, Силовые Тренировки, Кардиотренировки\n- Образ жизни: Умеренно активный\n- Ежедневное потребление калорий: 1800 ккал\n- Доступность тренировок: 4 дня в неделю\n- Среднее количество сна: 7 часов\n- Хронические заболевания отсутствуют\n- Травмы: Травма колена\n\n### Рекомендации по Питанию:\n- **Основные принципы питания**:\n  1. Придерживайтесь сбалансированной диеты, включающей все основные группы продуктов.\n  2. Обеспечьте достаточное количество белка для поддержания мышечной массы (примерно 1,6-2 г/кг массы тела).\n  3. Контролируйте общее количество потребляемых калорий, чтобы обеспечить дефицит для потери веса.\n  4. Избегайте чрезмерного потребления насыщенных жиров и трансжиров.\n  5. Пейте достаточно воды в течение дня.\n\n- **Меню на неделю**:\n    - Завтрак: Омлет с овощами и цельнозерновой тост\n    - Перекус: Греческий йогурт с фруктами\n    - Обед: Куриный салат с авокадо и зеленью\n    - Полдник: Горсть орехов или семян\n    - Ужин: Рыба на гриле с овощами\n    - Перед сном: Протеиновый коктейль или творог\n\n### Рекомендации по Физической Активности:\n- **Силовые Тренировки**:\n  1. Выполняйте силовые упражнения 2-3 раза в неделю, уделяя особое внимание ногам и корпусу (например, приседания, выпады, тяги, жимы лежа).\n  2. Используйте базовые многосуставные упражнения для максимального эффекта.\n  3. Включите суперсеты и дропсеты для увеличения интенсивности тренировок.\n\n- **Кардиотренировки**:\n  1. Включайте кардиоупражнения средней интенсивности 2-3 раза в неделю (например, бег трусцой, плавание, велотренажер).\n  2. Старайтесь выполнять кардио после силовой тренировки для повышения эффективности сжигания жира.\n\n- **Йога**:\n  1. Практикуйте йогу 1-2 раза в неделю для улучшения гибкости и снятия стресса.\n  2. Добавьте дыхательные практики и медитацию для расслабления и восстановления.\n\n- **Восстановление**:\n  1. Обеспечьте достаточный отдых между тренировками, особенно если у вас травма колена.\n  2. Следите за сигналами своего организма и не перегружайте себя.\n\n### Примерная Программа Тренировок на Неделю:\nПонедельник: Силовая тренировка ног и корпуса\nВторник: Кардио (бег трусцой)\nСреда: Йога\nЧетверг: Силовая тренировка верхней части тела\nПятница: Кардио (велотренажер)\nСуббота: День отдыха или легкая растяжка\nВоскресенье: Силовая тренировка всего тела\n\nЭтот план является лишь ориентировочным и может быть адаптирован под ваши индивидуальные потребности и предпочтения. Если у вас есть какие-либо вопросы или нужна дополнительная помощь, пожалуйста, обращайтесь!"
}
```

</details>