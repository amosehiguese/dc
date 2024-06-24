package store

import corerepo "github.com/amosehiguese/dc/pkg/core-repo"

type Attachment struct {
	corerepo.BaseModel
	MessageID string `gorm:"index;not null" json:"message_id"`
	SecureURL string `gorm:"not null" json:"secure_url"`
	PublicID  string `gorm:"not null" json:"public_id"`
}
