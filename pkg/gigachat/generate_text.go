package gigachat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	urlGenerateText = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
)

type RequestBody struct {
	Model           string     `json:"model"`
	Stream          bool       `json:"stream"`
	Update_interval int        `json:"update_interval"`
	Messages        []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionResult описывает результат генерации текста моделью
type ChatCompletionResult struct {
	Choices []Choice         `json:"choices"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Object  string           `json:"object"`
	Usage   UsageInformation `json:"usage"`
}

// Choice представляет один выбор модели
type Choice struct {
	Message      MessageResult `json:"message"`
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
}

// Message содержит контент и роль
type MessageResult struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// UsageInformation содержит информацию об использовании токенов
type UsageInformation struct {
	PromptTokens          int `json:"prompt_tokens"`
	CompletionTokens      int `json:"completion_tokens"`
	TotalTokens           int `json:"total_tokens"`
	PreCachedPromptTokens int `json:"pre_cached_prompt_tokens"`
}

func (gc *Client) GenerateText(body RequestBody) (ChatCompletionResult, error) {
	//
	var chatResult ChatCompletionResult

	//
	if gc.accessToken == "" {
		if err := gc.getAccessToken("GIGACHAT_API_PERS"); err != nil {
			return chatResult, err
		}
	}

	// Определение URL-адреса конечной точки
	urlEndpoint := urlGenerateText

	// Преобразуем структуру в JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		//fmt.Println("Ошибка маршалинга:", err)
		return chatResult, err
	}

	// Создание нового HTTP-запроса
	req, err := http.NewRequest("POST", urlEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		//log.Fatalf("Ошибка создания HTTP-запроса: %v", err)
		return chatResult, err
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.accessToken)) //
	//req.Header.Set("RqUID", "92d59172-a445-4ca5-bf59-7c986eec7f56")

	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	// c.client = &http.Client{Transport: tr}

	// Выполнение HTTP-запроса
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		//log.Fatalf("Ошибка выполнения HTTP-запроса: %v", err)
		return chatResult, err
	}
	defer resp.Body.Close()

	//
	var buf bytes.Buffer

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		//logger.Error(err)
		return chatResult, err
	}

	err = json.Unmarshal(buf.Bytes(), &chatResult)
	if err != nil {
		return chatResult, err
		//log.Fatalf("Ошибка: %v", err)
		//return
	}

	//
	return chatResult, nil //responseBody
}
