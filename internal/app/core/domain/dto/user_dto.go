package dto

type FindUserDTO struct {
	Id       interface{}  `json:"id"`
	Email    string       `json:"email,omitempty"`
	UserType string       `json:"user_type,omitempty"`
	Customer *CustomerDTO `json:"customer,omitempty"`
}

type LoginDTO struct {
	AccessToken string `json:"bearer_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
