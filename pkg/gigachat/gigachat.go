package gigachat

type GigaChatClient struct {
	// baseURL string
	apiKey      string
	accessToken string
}

func NewGigaChatClient(apiKey string) *GigaChatClient {
	return &GigaChatClient{
		// baseURL: baseURL,
		apiKey: apiKey,
	}
}
