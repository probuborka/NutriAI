package gigachat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

const (
	urlAccessToken = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

func (gc *GigaChatClient) getAccessToken(scope string) error {
	// Определение URL-адреса конечной точки
	urlEndpoint := urlAccessToken
	uuid := uuid.New().String()

	// Параметры для тела запроса
	data := url.Values{}
	data.Set("scope", scope) //"GIGACHAT_API_PERS"

	// Преобразование параметров в строку формата x-www-form-urlencoded
	body := bytes.NewBufferString(data.Encode())

	// Создание нового HTTP-запроса
	req, err := http.NewRequest("POST", urlEndpoint, body)
	if err != nil {
		return err
		//log.Fatalf("Ошибка создания HTTP-запроса: %v", err)
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.apiKey)) // ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOmEwN2Q5YjhkLWVlNDAtNDUzZS04MTk1LTYzMDQxODU0NjYwMA==
	req.Header.Set("RqUID", uuid)

	// Выполнение HTTP-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
		//log.Fatalf("Ошибка выполнения HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
		//log.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	//
	var accessToken AccessToken

	err = json.Unmarshal(responseBody, &accessToken)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		return err
		//log.Fatalf("Ошибка: %v", err)
		//return
	}

	gc.accessToken = accessToken.AccessToken

	// Вывод результата
	//fmt.Printf("Статус-код: %d\n", resp.StatusCode)
	//fmt.Printf("Тело ответа: %s\n", string(responseBody))
	return nil
}
