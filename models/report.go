package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	ID            string         `json:"id" gorm:"type:uuid;primaryKey"`
	Range         string         `json:"range"` // 1d, 7d, 30d
	TotalOrders   int64          `json:"total_orders"`
	TotalRevenue  float64        `json:"total_revenue"`
	StartDate     time.Time      `json:"start_date"`
	EndDate       time.Time      `json:"end_date"`
	CreatedAt     time.Time      `json:"created_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()
	return nil
}