package controller

import (
	"github.com/yotsaphon08/sa-64-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /register
func CreateRegister(c *gin.Context) {

	var register entity.Register
	var assessmentsheet entity.AssessmentSheet
	var ambulance entity.Ambulance
	var cases entity.Case
	var owner entity.User

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา cases ด้วย id
	if tx := entity.DB().Where("id = ?", register.CaseID).First(&cases); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cases not found"})
		return
	}

	// 10: ค้นหา assessmentsheet ด้วย id
	if tx := entity.DB().Where("id = ?", register.AssessmentSheetID).First(&assessmentsheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessmentsheet not found"})
		return
	}

	// 11: ค้นหา ambulance ด้วย id
	if tx := entity.DB().Where("id = ?", register.AmbulanceID).First(&ambulance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}
	// 12: ค้นหา owner ด้วย id
	if tx := entity.DB().Where("id = ?", register.CaseID).First(&owner); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner not found"})
		return
	}
	// 13: สร้าง Register
	rt := entity.Register{

		AssessmentSheet: assessmentsheet,       // โยงความสัมพันธ์กับ Entity AssessmentSheet
		Case:            cases,                 // โยงความสัมพันธ์กับ Entity Notify
		Ambulance:       ambulance,             // โยงความสัมพันธ์กับ Entity Ambulance
		RegisterTime:    register.RegisterTime, // ตั้งค่าฟิลด์ RegisterTime
		Owner:           owner,                 // โยงความสัมพันธ์กับ Entity User
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rt})
}

// GET /register/:id

func GetRegister(c *gin.Context) {

	var register entity.Register

	id := c.Param("id")

	if err := entity.DB().Preload("AssessmentSheet").Preload("Ambulance").Preload("Case").Preload("Owner").Raw("SELECT * FROM registers WHERE id = ?", id).Find(&register).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": register})

}

// GET /registers

func ListRegisters(c *gin.Context) {

	var registers []entity.Register

	if err := entity.DB().Preload("AssessmentSheet").Preload("Ambulance").Preload("Case").Preload("Owner").Raw("SELECT * FROM registers").Find(&registers).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": registers})

}

// DELETE /registers/:id

func DeleteRegister(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM registers WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "register not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /registers

func UpdateRegister(c *gin.Context) {

	var register entity.Register

	if err := c.ShouldBindJSON(&register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", register.ID).First(&register); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "register not found"})

		return

	}

	if err := entity.DB().Save(&register).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": register})

}
