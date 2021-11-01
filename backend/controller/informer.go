package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
	"golang.org/x/crypto/bcrypt"
)

// GET /informers
// List all informers
func ListInformer(c *gin.Context) {
	var informers []entity.Informer
	if err := entity.DB().Raw("SELECT * FROM informers").Scan(&informers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": informers})
}

// GET /informers/:id
// Get informers by id
func GetInformer(c *gin.Context) {
	var informers entity.Informer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM informers WHERE id = ?", id).Find(&informers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": informers})
}

// POST /informers
func CreateInformer(c *gin.Context) {
	var informers entity.Informer
	if err := c.ShouldBindJSON(&informers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(informers.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	informers.Password = string(bytes)

	if err := entity.DB().Create(&informers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": informers})
}

// PATCH /informers
func UpdateInformer(c *gin.Context) {
	var informers entity.Informer
	if err := c.ShouldBindJSON(&informers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", informers.ID).First(&informers); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "informers not found"})
		return
	}

	if err := entity.DB().Save(&informers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": informers})
}

// DELETE /informers/:id
func DeleteInformer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM informers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "informers not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Informer{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
