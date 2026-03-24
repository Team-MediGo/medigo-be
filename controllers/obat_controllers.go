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
	// ambil data form
	nama := c.PostForm("nama")
	kategori := c.PostForm("kategori")
	harga, _ := strconv.ParseFloat(c.PostForm("harga"), 64)
	stok, _ := strconv.Atoi(c.PostForm("stok"))

	//ambil file gambar
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//buka file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file"})
		return
	}
	defer src.Close()

	//upload ke Cloudinary
	imageURL, err := config.UploadImage(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal upload gambar: " + err.Error()})
		return
	}

	//simpan ke database
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

	// Ambil data dari form
	if nama := c.PostForm("nama"); nama != "" {
		obat.Nama = nama
	}
	if kategori := c.PostForm("kategori"); kategori != "" {
		obat.Kategori = kategori
	}
	if harga := c.PostForm("harga"); harga != "" {
		if h, err := strconv.ParseFloat(harga, 64); err == nil {
			obat.Harga = h
		}
	}
	if stok := c.PostForm("stok"); stok != "" {
		if s, err := strconv.Atoi(stok); err == nil {
			obat.Stok = s
		}
	}

	// Ganti gambar kalau ada file baru
	file, err := c.FormFile("image")
	if err == nil {
		src, err := file.Open()
		if err == nil {
			defer src.Close()
			imageURL, err := config.UploadImage(src)
			if err == nil {
				obat.ImageURL = imageURL
			}
		}
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
