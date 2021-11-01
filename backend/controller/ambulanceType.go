package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /AmbulanceType
func CreateAmbulanceType(c *gin.Context) {
	var ambulance_type entity.AmbulanceType
	if err := c.ShouldBindJSON(&ambulance_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&ambulance_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance_type})
}

// GET /Ambulancetype/:id
func GetAmbulanceType(c *gin.Context) {
	var ambulance_type entity.AmbulanceType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM ambulance_types WHERE id = ?", id).Scan(&ambulance_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance_type})
}

// GET /AmbulanceTypes
func ListAmbulanceType(c *gin.Context) {
	var ambulance_types []entity.AmbulanceType
	if err := entity.DB().Raw("SELECT * FROM ambulance_types").Scan(&ambulance_types).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance_types})
}

// DELETE /ambulanceTypes/:id
func DeleteAmbulanceType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulance_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ambulanceTypes
func UpdateAmbulanceType(c *gin.Context) {
	var ambulance_type entity.AmbulanceType
	if err := c.ShouldBindJSON(&ambulance_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", ambulance_type.ID).First(&ambulance_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}
	if err := entity.DB().Save(&ambulance_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance_type})
}
