package dto

type RequestLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
