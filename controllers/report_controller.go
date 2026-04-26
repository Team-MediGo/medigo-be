package controllers

import (
	"medigo-be/config"
	"medigo-be/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 🔥 Generate + Simpan Report
func GenerateAndSaveReport(c *gin.Context) {
	rangeParam := c.Query("range")

	var duration time.Duration

	switch rangeParam {
	case "1d":
		duration = 24 * time.Hour
	case "7d":
		duration = 7 * 24 * time.Hour
	case "30d":
		duration = 30 * 24 * time.Hour
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "range harus 1d, 7d, atau 30d",
		})
		return
	}

	startDate := time.Now().Add(-duration)
	endDate := time.Now()

	var totalOrders int64
	var totalRevenue float64

	db := config.DB

	// 🔥 hanya yang sudah dibayar
	query := db.Model(&models.Order{}).
		Where("created_at >= ? AND payment_status = ?", startDate, "paid")

	query.Count(&totalOrders)
	query.Select("COALESCE(SUM(total_harga), 0)").Scan(&totalRevenue)

	// Simpan ke DB
	report := models.Report{
		Range:        rangeParam,
		TotalOrders:  totalOrders,
		TotalRevenue: totalRevenue,
		StartDate:    startDate,
		EndDate:      endDate,
	}

	if err := db.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal menyimpan report",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "report berhasil disimpan",
		"data":    report,
	})
}

// Ambil Semua Report
func GetSavedReports(c *gin.Context) {
	var reports []models.Report

	if err := config.DB.Order("created_at desc").Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal mengambil data report",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": reports,
	})
}
