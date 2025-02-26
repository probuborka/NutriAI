package gigachat

type Client struct {
	// baseURL string
	apiKey      string
	accessToken string
}

func New(apiKey string) *Client {
	return &Client{
		// baseURL: baseURL,
		apiKey: apiKey,
	}
}
