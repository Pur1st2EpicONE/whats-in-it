package models

type Token interface {
	GetToken() string
}

type Response interface {
	GetResponse() (string, error)
}
