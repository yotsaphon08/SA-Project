package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/entity"
)

// POST /cases
func CreateCase(c *gin.Context) {

	var cases entity.Case
	var characteristics entity.Characteristic
	var patients entity.Patient
	var levels entity.Level
	var informers entity.Informer

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร case
	if err := c.ShouldBindJSON(&cases); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา characteristic ด้วย id
	if tx := entity.DB().Where("id = ?", cases.CharacteristicID).First(&characteristics); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "characteristics not found"})
		return
	}

	// 10: ค้นหา level ด้วย id
	if tx := entity.DB().Where("id = ?", cases.LevelID).First(&levels); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "levels not found"})
		return
	}

	// 11: ค้นหา patient ด้วย id
	if tx := entity.DB().Where("id = ?", cases.PatientID).First(&patients); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patients not found"})
		return
	}

	// 12: ค้นหา informer ด้วย id
	if tx := entity.DB().Where("id = ?", cases.InformerID).First(&informers); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "informers not found"})
		return
	}

	// 13: สร้าง Case
	cs := entity.Case{
		Characteristic: characteristics, // โยงความสัมพันธ์กับ Entity Characteristic
		Level:          levels,          // โยงความสัมพันธ์กับ Entity Level
		Patient:        patients,        // โยงความสัมพันธ์กับ Entity patient
		Informer:       informers,       // โยงความสัมพันธ์กับ Entity Informers
		CaseTime:       cases.CaseTime,  // ตั้งค่าฟิลด์ CaseTime
		Address:        cases.Address,   // ตั้งค่าฟิลด์ Address
	}

	// 14: บันทึก
	if err := entity.DB().Create(&cs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cs})
}

// GET /case/:id
func GetCase(c *gin.Context) {
	var cases entity.Case
	id := c.Param("id")
	if err := entity.DB().Preload("Characteristic").Preload("Level").Preload("Patient").Preload("Informer").Raw("SELECT * FROM cases WHERE id = ?", id).Find(&cases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cases})
}

// GET /cases
func ListCase(c *gin.Context) {
	var cases []entity.Case
	if err := entity.DB().Preload("Characteristic").Preload("Level").Preload("Patient").Preload("Informer").Raw("SELECT * FROM cases").Find(&cases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cases})
}

// DELETE /cases/:id
func DeleteCase(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cases WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "watchvideo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /cases
func UpdateCase(c *gin.Context) {
	var cases entity.Case
	if err := c.ShouldBindJSON(&cases); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cases.ID).First(&cases); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cases not found"})
		return
	}

	if err := entity.DB().Save(&cases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cases})
}
