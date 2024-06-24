package store

import corerepo "github.com/amosehiguese/dc/pkg/core-repo"

type Reaction struct {
	corerepo.BaseModel
	MessageID string `gorm:"index;not null" json:"message_id"`
	UserID    string `gorm:"index;not null" json:"user_id"`
	Emoji     string `gorm:"not null" json:"emoji,omitempty"`
}
