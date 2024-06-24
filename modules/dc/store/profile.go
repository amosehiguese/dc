package store

import (
	corejsonb "github.com/amosehiguese/dc/pkg/core-jsonb"
	corerepo "github.com/amosehiguese/dc/pkg/core-repo"
)

type Profile struct {
	corerepo.BaseModel
	AvatarURL      string          `gorm:"not null" json:"avatar_url" validate:"url"`
	AvatarPublicID string          `gorm:"not null" json:"avatar_public_id" validate:"url"`
	City           string          `gorm:"size:255;not null" json:"city" validate:"lte=255"`
	State          string          `gorm:"size:255;not null" json:"state" validate:"lte=255"`
	Country        string          `gorm:"size:255;not null" json:"country" validate:"lte=255"`
	Bio            string          `gorm:"size:255;not null" json:"bio" validate:"lte=255"`
	MedicalHistory corejsonb.JSONB `json:"medical_history"`
}

func (u *Profile) TableName() string {
	return "profiles"
}
