package controller

import (
	"github.com/yotsaphon08/sa-64-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateCarPath(c *gin.Context) {

	var car_path entity.Car_path

	if err := c.ShouldBindJSON(&car_path); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&car_path).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": car_path})

}

// GET /user/:id

func GetCarPath(c *gin.Context) {

	var car_path entity.Car_path

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM car_paths WHERE id = ?", id).Scan(&car_path).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": car_path})

}

// GET /users

func ListCarPath(c *gin.Context) {

	var car_path []entity.Car_path

	if err := entity.DB().Raw("SELECT * FROM car_paths").Scan(&car_path).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": car_path})

}

// DELETE /users/:id

func DeleteCarPath(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM car_paths WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "car_path not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateCarPath(c *gin.Context) {

	var car_path entity.Car_path

	if err := c.ShouldBindJSON(&car_path); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", car_path.ID).First(&car_path); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "car_path not found"})

		return

	}

	if err := entity.DB().Save(&car_path).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": car_path})

}
