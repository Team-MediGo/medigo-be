package controllers

import (
	"medigo-be/config"
	"medigo-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /cart
func GetCart(c *gin.Context) {
	userID := c.GetString("user_id")

	var items []models.Cart
	config.DB.Preload("Obat").Where("user_id = ?", userID).Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

// POST /cart
func AddToCart(c *gin.Context) {
	userID := c.GetString("user_id")

	var body struct {
		ObatID   string `json:"obat_id"`
		Quantity int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah obat sudah ada di cart
	var existing models.Cart
	err := config.DB.Where("user_id = ? AND obat_id = ?", userID, body.ObatID).First(&existing).Error

	if err == nil {
		// Sudah ada → update quantity
		existing.Quantity += body.Quantity
		config.DB.Save(&existing)
		c.JSON(http.StatusOK, gin.H{"data": existing, "message": "Quantity diupdate"})
		return
	}

	// Belum ada → tambah baru
	cart := models.Cart{
		UserID:   userID,
		ObatID:   body.ObatID,
		Quantity: body.Quantity,
	}
	config.DB.Create(&cart)
	config.DB.Preload("Obat").First(&cart, "id = ?", cart.ID)
	c.JSON(http.StatusCreated, gin.H{"data": cart})
}

// PUT /cart/:id
func UpdateCart(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	var cart models.Cart
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}

	var body struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Quantity <= 0 {
		config.DB.Delete(&cart)
		c.JSON(http.StatusOK, gin.H{"message": "Item dihapus dari cart"})
		return
	}

	cart.Quantity = body.Quantity
	config.DB.Save(&cart)
	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DELETE /cart/:id
func DeleteCart(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	var cart models.Cart
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
		return
	}

	config.DB.Delete(&cart)
	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil dihapus"})
}

// DELETE /cart
func ClearCart(c *gin.Context) {
	userID := c.GetString("user_id")
	config.DB.Where("user_id = ?", userID).Delete(&models.Cart{})
	c.JSON(http.StatusOK, gin.H{"message": "Cart berhasil dikosongkan"})
}
