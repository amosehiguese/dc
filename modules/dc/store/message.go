package store

import corerepo "github.com/amosehiguese/dc/pkg/core-repo"

type Message struct {
	corerepo.BaseModel
	ChatID       string        `gorm:"index;not null" json:"chat_id"`
	SenderID     string        `gorm:"index;not null" json:"sender_id"`
	Content      string        `gorm:"not null" json:"content"`
	Attachments  []*Attachment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"attachments,omitempty"`
	AttachmentID string
	URL          string      `json:"url,omitempty"`
	Reactions    []*Reaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"reactions,omitempty"`
	ReactionID   string
}
