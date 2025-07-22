package models

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

type ChatRequest struct {
	Model          string    `json:"model"`
	Stream         bool      `json:"stream"`
	UpdateInterval int       `json:"update_interval"`
	Messages       []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
