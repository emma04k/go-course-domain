package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primaryKey;uniqueIndex"`
	Name      string         `json:"name" gorm:"type:char(50); not null"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAT gorm.DeletedAt `json:"-"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Id == "" {
		c.Id = uuid.NewString()
	}
	return
}
