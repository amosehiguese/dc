package store

import corerepo "github.com/amosehiguese/dc/pkg/core-repo"

type Doctor struct {
	corerepo.BaseModel
	UserID         string `gorm:"index;not null" json:"user_id"`
	Specialization string `gorm:"not null" json:"specialization"`
	Experience     string `gorm:"not null" json:"experience"`
}
