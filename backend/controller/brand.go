package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /brands
func CreateBrand(c *gin.Context) {
	var brand entity.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// GET /brand/:id
func GetBrand(c *gin.Context) {
	var brand entity.Brand
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM brands WHERE id = ?", id).Scan(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": brand})
}

// GET /brands
func ListBrand(c *gin.Context) {
	var brands []entity.Brand
	if err := entity.DB().Raw("SELECT * FROM brands").Scan(&brands).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": brands})
}

// DELETE /brands/:id
func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM brands WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /brands
func UpdateBrand(c *gin.Context) {
	var brand entity.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", brand.ID).First(&brand); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand not found"})
		return
	}
	if err := entity.DB().Save(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": brand})
}
