package store

import corerepo "github.com/amosehiguese/dc/pkg/core-repo"

type Connection struct {
	corerepo.BaseModel
	UserID   string `gorm:"index;not null" json:"user_id"`
	DoctorID string `gorm:"index;not null" json:"doctor_id"`
}
