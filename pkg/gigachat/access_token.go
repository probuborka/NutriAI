package gigachat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

const (
	urlAccessToken = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

func (gc *Client) getAccessToken(scope string) error {
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.apiKey))
	req.Header.Set("RqUID", uuid)

	// Выполнение HTTP-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
		//log.Fatalf("Ошибка выполнения HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	//
	var buf bytes.Buffer

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	//
	var accessToken AccessToken

	err = json.Unmarshal(buf.Bytes(), &accessToken)
	if err != nil {
		return err
	}

	gc.accessToken = accessToken.AccessToken

	return nil
}
