package model

type Client struct {
	ID           int    `json:"id"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	CallbackURL  string `json:"callback_url"`
	Status       string `json:"status"`
}
