package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" grom:"type:uuid;default:gen_random_uuid();primaryKey"`
	Nama      string         `json:"nama"`
	Email     string         `jason:"email" gorm:"unique"`
	Password  string         `json:"-"`
	Role      string         `json:"role" gorm:"dafault:admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
