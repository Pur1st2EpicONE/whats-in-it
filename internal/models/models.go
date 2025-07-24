package models

type Token interface {
	GetToken() string
}

type GigaChatToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

func (g GigaChatToken) GetToken() string {
	return g.AccessToken
}

type GigaChatRequest struct {
	Model          string            `json:"model"`
	Stream         bool              `json:"stream"`
	UpdateInterval int               `json:"update_interval"`
	Messages       []GigaChatMessage `json:"messages"`
}

type GigaChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response interface {
	GetResponse() string
}

type GigaChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (g GigaChatResponse) GetResponse() string {
	return g.Choices[0].Message.Content
}

type YandexToken struct {
	AccessToken string `json:"iamToken"`
	ExpiresAt   int    `json:"expiresAt"`
}

func (i YandexToken) GetToken() string {
	return i.AccessToken
}

type YandexResponse struct {
	Result struct {
		Alternatives []struct {
			Message struct {
				Role string `json:"role"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"alternatives"`
	} `json:"result"`
}

func (y YandexResponse) GetResponse() string {
	return y.Result.Alternatives[0].Message.Text
}

type YandexTokenRequest struct {
	OauthToken string `json:"yandexPassportOauthToken"`
}
