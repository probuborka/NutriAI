package gigachat

type Client struct {
	apiKey      string
	accessToken string
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
