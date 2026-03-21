package controllers

import (
	"medigo-be/config"
	"medigo-be/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /obat
func GetAllObat(c *gin.Context) {
	var obat []models.Obat
	config.DB.Find(&obat)
	c.JSON(http.StatusOK, gin.H{"data": obat})
}

// GET /obat/:id
func GetObatByID(c *gin.Context) {
	var obat models.Obat
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&obat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": obat})
}

// POST /obat
func CreateObat(c *gin.Context) {
	// Ambil data form
	nama := c.PostForm("nama")
	kategori := c.PostForm("kategori")
	harga, _ := strconv.ParseFloat(c.PostForm("harga"), 64)
	stok, _ := strconv.Atoi(c.PostForm("stok"))

	// Ambil file gambar
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar wajib diupload"})
		return
	}

	// Buka file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file"})
		return
	}
	defer src.Close()

	// Upload ke Cloudinary
	imageURL, err := config.UploadImage(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal upload gambar: " + err.Error()})
		return
	}

	// Simpan ke database
	obat := models.Obat{
		Nama:     nama,
		Kategori: kategori,
		Harga:    harga,
		Stok:     stok,
		ImageURL: imageURL,
	}

	result := config.DB.Create(&obat)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": obat})
}

// PUT /obat/:id
func UpdateObat(c *gin.Context) {
	var obat models.Obat
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&obat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&obat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&obat)
	c.JSON(http.StatusOK, gin.H{"data": obat})
}

// DELETE /obat/:id
func DeleteObat(c *gin.Context) {
	var obat models.Obat
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&obat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat tidak ditemukan"})
		return
	}
	config.DB.Delete(&obat)
	c.JSON(http.StatusOK, gin.H{"message": "Obat berhasil dihapus"})
}
