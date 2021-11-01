package controller

import (
	"github.com/yotsaphon08/sa-64-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /assess

func CreateAssess(c *gin.Context) {

	var assess entity.Assess

	if err := c.ShouldBindJSON(&assess); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// GET /assess/:id

func GetAssess(c *gin.Context) {

	var assess entity.Assess

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM assesses WHERE id = ?", id).Scan(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// GET /assess

func ListAssess(c *gin.Context) {

	var assess []entity.Assess

	if err := entity.DB().Raw("SELECT * FROM assesses").Scan(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}

// DELETE /assess/:id

func DeleteAssess(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM assesses WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "assesses not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /assess

func UpdateAssess(c *gin.Context) {

	var assess entity.Assess

	if err := c.ShouldBindJSON(&assess); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", assess.ID).First(&assess); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "assesses not found"})

		return

	}

	if err := entity.DB().Save(&assess).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": assess})

}
