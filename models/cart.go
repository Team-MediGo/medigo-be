package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string         `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    string         `json:"user_id" gorm:"type:uuid"`
	ObatID    string         `json:"obat_id" gorm:"type:uuid"`
	Quantity  int            `json:"quantity"`
	Obat      Obat           `json:"obat" gorm:"foreignKey:ObatID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}
