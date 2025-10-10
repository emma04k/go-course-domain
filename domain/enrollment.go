package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primaryKey;uniqueIndex"`
	UserId    string  `json:"user_id,omitempty" gorm:"type:char(36)"`
	User      *User   `json:"user,omitempty"`
	CourseId  string  `json:"course_id" gorm:"type:char(36);not null"`
	Course    *Course `json:"course,omitempty"`
	Status    string  `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAT gorm.DeletedAt `json:"-"`
}

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Id == "" {
		c.Id = uuid.NewString()
	}
	return
}
