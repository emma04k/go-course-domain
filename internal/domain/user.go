package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primaryKey;uniqueIndex"`
	FirstName string         `json:"first_name" gorm:"type:char(50); not null"`
	LastName  string         `json:"last_name" gorm:"type:char(50); not null"`
	Email     string         `json:"email" gorm:"type:char(100); not null"`
	Phone     string         `json:"phone" gorm:"type:char(20); not null"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAT gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.NewString()
	}
	return
}
