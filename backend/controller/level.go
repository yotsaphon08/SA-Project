package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /levels
func CreateLevel(c *gin.Context) {
	var levels entity.Level
	if err := c.ShouldBindJSON(&levels); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": levels})
}

// GET /level/:id
func GetLevel(c *gin.Context) {
	var levels entity.Level
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM levels WHERE id = ?", id).Scan(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": levels})
}

// GET /levels
func ListLevels(c *gin.Context) {
	var levels []entity.Level
	if err := entity.DB().Raw("SELECT * FROM levels").Scan(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": levels})
}

// DELETE /levels/:id
func DeleteLevel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM levels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "levels not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /levels
func UpdateLevel(c *gin.Context) {
	var levels entity.Level
	if err := c.ShouldBindJSON(&levels); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", levels.ID).First(&levels); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "levels not found"})
		return
	}

	if err := entity.DB().Save(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": levels})
}
