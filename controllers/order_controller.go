package controllers

import (
	"medigo-be/config"
	"medigo-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /orders — checkout dari cart
func CreateOrder(c *gin.Context) {
	userID := c.GetString("user_id")

	var body struct {
		AlamatAntar string `json:"alamat_antar"`
		MetodeBayar string `json:"metode_bayar"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil semua item di cart
	var cartItems []models.Cart
	config.DB.Preload("Obat").Where("user_id = ?", userID).Find(&cartItems)

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart kosong"})
		return
	}

	// Hitung total harga
	var totalHarga float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		subtotal := item.Obat.Harga * float64(item.Quantity)
		totalHarga += subtotal
		orderItems = append(orderItems, models.OrderItem{
			ObatID:   item.ObatID,
			NamaObat: item.Obat.Nama,
			Harga:    item.Obat.Harga,
			Quantity: item.Quantity,
			Subtotal: subtotal,
			ImageURL: item.Obat.ImageURL,
		})
	}

	// Buat order
	order := models.Order{
		UserID:      userID,
		TotalHarga:  totalHarga,
		AlamatAntar: body.AlamatAntar,
		MetodeBayar: body.MetodeBayar,
		Status:      "pending",
		Items:       orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat order"})
		return
	}

	// Kosongkan cart setelah checkout
	config.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order berhasil dibuat",
		"data":    order,
	})
}

// GET /orders — riwayat order user
func GetOrders(c *gin.Context) {
	userID := c.GetString("user_id")

	var orders []models.Order
	config.DB.Preload("Items").Where("user_id = ?", userID).Order("created_at desc").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GET /orders/:id — detail order
func GetOrderByID(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	var order models.Order
	if err := config.DB.Preload("Items").Where("id = ? AND user_id = ?", id, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// PUT /orders/:id/status — update status (untuk admin/kurir)
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}

	order.Status = body.Status
	config.DB.Save(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}
