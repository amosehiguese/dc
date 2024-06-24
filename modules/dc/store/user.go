package store

import (
	"time"

	coreauth "github.com/amosehiguese/dc/pkg/core-auth"
	corerepo "github.com/amosehiguese/dc/pkg/core-repo"
	coreutils "github.com/amosehiguese/dc/pkg/core-utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	corerepo.BaseModel
	FirstName                   string        `gorm:"size:255;not null" json:"firstname" validate:"required"`
	LastName                    *string       `gorm:"size:255" json:"lastname,omitempty" validate:"lte=255"`
	Email                       string        `gorm:"uniqueIndex;not null" json:"email" validate:"required,email,lte=255"`
	Password                    string        `gorm:"size:255;not null" json:"-" validate:"required,lte=255"`
	Role                        coreauth.Role `gorm:"size:25;not null" json:"role" validate:"required,lte=25"`
	Verified                    *time.Time    `json:"verified,omitempty"`
	IsVerified                  bool          `gorm:"default:false" json:"is_verified"`
	VerificationToken           string        `gorm:"size:255" json:"-" validate:"required,lte=255"`
	PasswordToken               *string       `gorm:"size:255" json:"-"`
	PasswordTokenExpirationDate *time.Time    `json:"password_token_expiration_date,omitempty"`
	NotificationEnabled         bool          `gorm:"default:false" json:"notification_enabled"`
	LastSeen                    time.Time     `json:"last_seen,omitempty"`
	Profile                     Profile       `json:"profile"`
	ProfileID                   string
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	hashedPwd := coreutils.HashPassword(u.Password)
	tx.Statement.SetColumn("password", hashedPwd)
	return nil
}

func (u *User) ComparePasswordHash(inputPwd string) bool {
	userPassword := coreutils.NormalizePassword(u.Password)
	inputPassword := coreutils.NormalizePassword(inputPwd)

	if err := bcrypt.CompareHashAndPassword(userPassword, inputPassword); err != nil {
		return false
	}
	return true

}
