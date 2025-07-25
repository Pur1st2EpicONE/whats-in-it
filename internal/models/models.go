package models

type Token interface {
	GetToken() string
}

type Response interface {
	GiveAnswer() (string, error)
}
