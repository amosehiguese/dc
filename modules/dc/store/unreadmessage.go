package store

import (
	"time"

	corerepo "github.com/amosehiguese/dc/pkg/core-repo"
)

type UnreadMessage struct {
	corerepo.BaseModel
	UserID    string    `gorm:"index;not null" json:"user_id"`
	ChatID    string    `gorm:"index;not null" json:"chat_id"`
	SenderID  string    `gorm:"index;not null" json:"sender_id"`
	MessageID string    `gorm:"index;not null" json:"message_id"`
	Count     int       `gorm:"default:1" json:"count"`
	ReadAt    time.Time `json:"read_at"`
}
