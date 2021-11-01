package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /ambulances
func CreateAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	var ambulanceType entity.AmbulanceType
	var status entity.Status
	var brand entity.Brand
	var owner entity.User

	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//  ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.ID).First(&owner); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 8: ค้นหา type ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.AmbulanceTypeID).First(&ambulanceType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}
	// 9: ค้นหา brand ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.BrandID).First(&brand); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand not found"})
		return
	}
	// 10: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}
	// 11: สร้าง Ambulance
	wv := entity.Ambulance{
		AmbulanceType: ambulanceType,           // โยงความสัมพันธ์กับ Entity AmbulanceTypeID
		Brand:         brand,                   // โยงความสัมพันธ์กับ Entity Brand
		Status:        status,                  // โยงความสัมพันธ์กับ Entity Status
		Owner:         owner,                   // โยงความสัมพันธ์กับ Entity User
		RecordingTime: ambulance.RecordingTime, // ตั้งค่าฟิลด์ RecordingTime
		Registration:  ambulance.Registration,  // ตั้งค่าฟิลด์ Registration
	}

	// 12: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /ambulance/:id
func GetAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	id := c.Param("id")
	if err := entity.DB().Preload("AmbulanceType").Preload("Brand").Preload("Status").Preload("Owner").Raw("SELECT * FROM ambulances WHERE id = ?", id).Find(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// GET /ambulances
func ListAmbulance(c *gin.Context) {
	var ambulances []entity.Ambulance
	if err := entity.DB().Preload("AmbulanceType").Preload("Brand").Preload("Status").Preload("Owner").Raw("SELECT * FROM ambulances").Find(&ambulances).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ambulances})
}

// DELETE /ambulances/:id
func DeleteAmbulance(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulances WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ambulances
func UpdateAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", ambulance.ID).First(&ambulance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}
	if err := entity.DB().Save(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}
