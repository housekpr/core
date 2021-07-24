// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Pump struct {
	State bool `json:"state"`
}

type PumpState struct {
	State bool `json:"state"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Tank struct {
	Level float64 `json:"level"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
