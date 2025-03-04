package gigachat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
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

var (
	ErrorBadRequest         = errors.New("bad request")
	ErrorAuthorizationError = errors.New("authorization error")
)

func (gc *Client) getAccessToken(scope string) error {
	//URL
	urlEndpoint := urlAccessToken

	//
	uuid := uuid.New().String()

	//body
	data := url.Values{}
	data.Set("scope", scope) //"GIGACHAT_API_PERS"

	//x-www-form-urlencoded
	body := bytes.NewBufferString(data.Encode())

	//http request
	req, err := http.NewRequest("POST", urlEndpoint, body)
	if err != nil {
		return err
	}

	//header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.apiKey))
	req.Header.Set("RqUID", uuid)

	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}

	//http client
	client := &http.Client{Transport: tr}

	//do
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return ErrorBadRequest
	case http.StatusUnauthorized:
		return ErrorAuthorizationError
	}

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
