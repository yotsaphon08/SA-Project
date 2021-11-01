package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /characteristics
func CreateCharacteristic(c *gin.Context) {
	var characteristics entity.Characteristic
	if err := c.ShouldBindJSON(&characteristics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&characteristics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": characteristics})
}

// GET /characteristic/:id
func GetCharacteristic(c *gin.Context) {
	var characteristics entity.Characteristic
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM characteristics WHERE id = ?", id).Scan(&characteristics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characteristics})
}

// GET /characteristics
func ListCharacteristics(c *gin.Context) {
	var characteristics []entity.Characteristic
	if err := entity.DB().Raw("SELECT * FROM characteristics").Scan(&characteristics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characteristics})
}

// DELETE /characteristics/:id
func DeleteCharacteristic(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM characteristics WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "characteristics not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /characteristics
func UpdateCharacteristic(c *gin.Context) {
	var characteristics entity.Characteristic
	if err := c.ShouldBindJSON(&characteristics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", characteristics.ID).First(&characteristics); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "characteristic not found"})
		return
	}

	if err := entity.DB().Save(&characteristics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characteristics})
}
