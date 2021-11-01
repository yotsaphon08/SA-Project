package controller

import (
	"github.com/yotsaphon08/sa-64-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateState(c *gin.Context) {

	var state entity.State

	if err := c.ShouldBindJSON(&state); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&state).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": state})

}

// GET /user/:id

func GetState(c *gin.Context) {

	var state entity.State

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM states WHERE id = ?", id).Scan(&state).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": state})

}

// GET /users

func ListState(c *gin.Context) {

	var state []entity.State

	if err := entity.DB().Raw("SELECT * FROM states").Scan(&state).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": state})

}

// DELETE /users/:id

func DeleteState(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM states WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "states not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateState(c *gin.Context) {

	var state entity.State

	if err := c.ShouldBindJSON(&state); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", state.ID).First(&state); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "states not found"})

		return

	}

	if err := entity.DB().Save(&state).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": state})

}
