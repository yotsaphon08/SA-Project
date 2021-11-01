package controller

import (
	"github.com/yotsaphon08/sa-64-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /users
// List all users
func ListPathStatus(c *gin.Context) {
	var path_status []entity.Path_status
	if err := entity.DB().Raw("SELECT * FROM path_statuses").Scan(&path_status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": path_status})
}

// GET /user/:id
// Get user by id
func GetPathStatus(c *gin.Context) {
	var path_status entity.Path_status
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM path_statuses WHERE id = ?", id).Scan(&path_status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": path_status})
}

// POST /users
func CreatePathStatus(c *gin.Context) {
	var path_status entity.Path_status
	if err := c.ShouldBindJSON(&path_status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&path_status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": path_status})
}

// PATCH /users
func UpdatePathStatus(c *gin.Context) {
	var path_status entity.Path_status
	if err := c.ShouldBindJSON(&path_status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", path_status.ID).First(&path_status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path_status not found"})
		return
	}

	if err := entity.DB().Save(&path_status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": path_status})
}

// DELETE /users/:id
func DeletePathStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM path_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path_status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
