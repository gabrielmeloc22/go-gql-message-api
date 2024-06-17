package model

import (
	"time"
)

type Chat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Members  []*User    `json:"members,omitempty" gorm:"many2many:user_chats"`
	Messages []*Message `json:"messages,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`

	FromId string `json:"from_id"`

	ChatId string `json:"chat_id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`

	Chats    []*Chat    `json:"chats" gorm:"many2many:user_chats"`
	Messages []*Message `json:"messages" gorm:"foreignKey:FromId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
