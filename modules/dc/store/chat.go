package store

import (
	"time"

	corerepo "github.com/amosehiguese/dc/pkg/core-repo"
)

type Chat struct {
	corerepo.BaseModel
	Name            string  `gorm:"size:255;not null" json:"name" validate:"lte=255"`
	IsGroupChat     bool    `gorm:"default:false" json:"is_group_chat"`
	Members         []*User `gorm:"many2many:chat_members;" json:"members" validate:"required,dive,min=1,max=10"`
	AvatarURL       string  `gorm:"not null" json:"avatar_url"`
	AvatarPublicID  string  `gorm:"not null" json:"avatar_public_id"`
	AdminID         string  `gorm:"index" json:"admin_id"`
	LatestMessageID string  `gorm:"index" json:"latest_message_id"`
}

type ChatMembers struct {
	ChatID    string `gorm:"primary_key"`
	UserID    string `gorm:"primary_key"`
	CreatedAt time.Time
}
