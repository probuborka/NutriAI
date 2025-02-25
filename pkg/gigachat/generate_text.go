package gigachat

import (
	"bytes"
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

func (gc *GigaChatClient) GenerateText(body RequestBody) (error, ChatCompletionResult) {

	//
	if gc.accessToken == "" {
		if err := gc.getAccessToken("GIGACHAT_API_PERS"); err != nil {
			return err, ChatCompletionResult{}
		}
	}

	// Определение URL-адреса конечной точки
	urlEndpoint := urlGenerateText

	// Преобразуем структуру в JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		//fmt.Println("Ошибка маршалинга:", err)
		return err, ChatCompletionResult{}
	}

	// Создание нового HTTP-запроса
	req, err := http.NewRequest("POST", urlEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		//log.Fatalf("Ошибка создания HTTP-запроса: %v", err)
		return err, ChatCompletionResult{}
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.accessToken)) // eyJjdHkiOiJqd3QiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwiYWxnIjoiUlNBLU9BRVAtMjU2In0.nLpO5s-VnQWLPJvv71SF09Z4cchcv78E62wFNGN6mEAr0RveclPyMXqyIsKw2mRgH1AodfeV9WKSR9qLmiIcBSat-ZXfnXIqmkUdops-GdUPJmVsU6lQBaIsBmEPamesmqNsFolq1UlyuGRWjAUXDSp0Uy2muBvHKhxtjLpQlqkJMS68qzEscBPUN7XCN3oEf-Kgw0KJZLG_nBG1tNyeEfr3T8GKkR0YDzgN7Kwy_T4QAIn9cPY7Z_B5CCgUuxazdUJ6XG054tmGBlW_mJpVlC9aZa56gdfajSHSB2gp6XVOZcLmQgk5xG1w-NhEAGXUGr5FSKzqPuNazbXkIV3IsQ.P3JHgydh4zN227VYiJ6daQ.WTULn1voe-3ed-ElozMsw8VTEIm9IDqFZ-kG00x1i5uu-_ReaF6qz40uH2sY3KByXaSX0dD6PapTTRWIHALSS4fJWSj5aWtcz5gcyWx3RpW8qBoW26mwb59UbOQpc2uRVcP9cM27YOrXpnutUtZdEUOUtYhVtosgadOq06JkMPBF32_9V58FmT2jPKFhLw-run_FcJCMNPO3GwDUPvSy4GceYBkbvOUy1r3YEA9ikcgfMgX6Rtth_Jhpc2XF1owL3t0uoWQUGlaJNbsrU1HRFXZZvWPZI7y53eVARTMWEURLaYoWWDm8idVj24JLdaPgZx6GWMnRu3TS5QyoZwfNI-wN5FM69NwBHegf92ul0dXSY4algq65WH_wOEAVE6ocYGSE1aoZmDmqRYmJto7cISJ0SedpKbpdNW6OMxJ42yKwoK44aTHAWli4ekF90MzcgQHLYicGHA2AYR32KFez_qJLfVyV3Kyt4Wr9yJLjkeKRCpi1-obO-ntHdGN8Cvtp3ob0FgIMF3ByVwYUogKcsIgdUJrsUw0krpUfjucuBdJnmh3Dzvnb42U-MR2Qd_D2R5NF3Mn282tBxz_zGwSJI9Uhc7DKhxjbO6MiEmltzfTpYDZPJScOqeCFfgu9Ybt3JYD363_DoiDR_5MSROk2MZ5dSbGL7LmBIlv9xt4HyULeRmO7-cPJs-4Mum-UhLOqj9c-PvDMHhLhRGHKxMBmGGO12j9wbrWKei1D-t9z35E.FL9r-tot2DFxk16RSz4tSYnqp7RBstni0LdkdzWBdQM")
	//req.Header.Set("RqUID", "92d59172-a445-4ca5-bf59-7c986eec7f56")

	// Выполнение HTTP-запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//log.Fatalf("Ошибка выполнения HTTP-запроса: %v", err)
		return err, ChatCompletionResult{}
	}
	defer resp.Body.Close()

	// Чтение ответа
	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	//log.Fatalf("Ошибка чтения тела ответа: %v", err)
	// 	return err, ChatCompletionResult{}
	// }

	//
	var buf bytes.Buffer

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		//logger.Error(err)
		return err, ChatCompletionResult{}
	}

	//
	var chatResult ChatCompletionResult

	err = json.Unmarshal(buf.Bytes(), &chatResult)
	if err != nil {
		//
		//response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		return err, ChatCompletionResult{}
		//log.Fatalf("Ошибка: %v", err)
		//return
	}

	//_ = responseBody

	// Вывод результата
	//fmt.Printf("Статус-код: %d\n", resp.StatusCode)
	//fmt.Printf("Тело ответа: %s\n", string(responseBody))

	//
	return nil, chatResult //responseBody
}
