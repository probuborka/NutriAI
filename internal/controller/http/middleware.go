package http

import (
	"context"
	"net/http"
	"time"

	"github.com/probuborka/NutriAI/internal/entity"
)

type middleware func(http.Handler) http.Handler

func compileMiddleware(h http.Handler, m []middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

type metric interface {
	RecordMetric(ctx context.Context, metric entity.Metric) error
}

func (h handler) RecordMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Выполняем следующий обработчик
		next.ServeHTTP(w, r)

		// Записываем метрики
		duration := time.Since(start).Seconds()

		// Метрика для количества запросов
		h.metric.RecordMetric(r.Context(), entity.Metric{
			Type:  entity.MetricTypeCounter,
			Name:  "http_requests_total",
			Value: 1,
			Labels: map[string]string{
				"method":   r.Method,
				"endpoint": r.URL.Path,
			},
		})

		// Метрика для времени обработки запроса
		h.metric.RecordMetric(r.Context(), entity.Metric{
			Type:  entity.MetricTypeHistogram,
			Name:  "http_request_duration_seconds",
			Value: duration,
			Labels: map[string]string{
				"method":   r.Method,
				"endpoint": r.URL.Path,
			},
		})

	})
}

// func logging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		logger.Infof("%s %s %s", req.Method, req.RequestURI, time.Now())
// 		next.ServeHTTP(w, req)
// 	})
// }

// func authentication(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// смотрим наличие пароля
// 		if len(cfg.Password) > 0 {
// 			var token string // JWT-токен из куки
// 			// получаем куку
// 			cookie, err := r.Cookie("token")
// 			if err == nil {
// 				token = cookie.Value
// 			}
// 			//
// 			secret := []byte(cfg.Password)
// 			// здесь код для валидации и проверки JWT-токена
// 			jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 				// секретный ключ для всех токенов одинаковый, поэтому просто возвращаем его
// 				return secret, nil
// 			})
// 			if err != nil {
// 				response(w, entityerror.Error{Error: err.Error()}, http.StatusUnauthorized)
// 				return
// 			}
// 			if !jwtToken.Valid {
// 				// возвращаем ошибку авторизации 401
// 				response(w, entityerror.Error{Error: "Authentification required"}, http.StatusUnauthorized)
// 				return
// 			}
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
