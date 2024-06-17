// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token string `json:"token"`
}

type Mutation struct {
}

type NewChatInput struct {
	Name  string `json:"name"`
	Users []*int `json:"users,omitempty"`
}

type NewMessageInput struct {
	Content   string    `json:"content"`
	To        int       `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
}

type Query struct {
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Token string `json:"token"`
}