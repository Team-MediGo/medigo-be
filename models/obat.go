package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Obat struct {
	ID        string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Nama      string         `json:"nama"`
	Kategori  string         `json:"kategori"`
	Harga     float64        `json:"harga"`
	Stok      int            `json:"stok"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (o *Obat) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
