package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID            string         `json:"id" gorm:"type:uuid;primaryKey"`
	UserID        string         `json:"user_id" gorm:"type:uuid"`
	Status        string         `json:"status" gorm:"default:pending"`
	TotalHarga    float64        `json:"total_harga"`
	AlamatAntar   string         `json:"alamat_antar"`
	MetodeBayar   string         `json:"metode_bayar"`
	PaymentStatus string         `json:"payment_status" gorm:"default:unpaid"`
	PaymentURL    string         `json:"payment_url"`
	Items         []OrderItem    `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type OrderItem struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	OrderID   string    `json:"order_id" gorm:"type:uuid"`
	ObatID    string    `json:"obat_id" gorm:"type:uuid"`
	NamaObat  string    `json:"nama_obat"`
	Harga     float64   `json:"harga"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	oi.ID = uuid.New().String()
	return nil
}
