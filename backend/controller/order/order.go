package controller

import (
	"net/http"

	"github.com/JusSix1/TwitterAccountDataBase/entity"
	"github.com/gin-gonic/gin"
)

// GET /order/:email
func GetLastAccont(c *gin.Context) {
	var user entity.User
	var order []entity.Order

	email := c.Param("email")

	if tx := entity.DB().Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := entity.DB().Raw("SELECT * FROM accounts WHERE user_id = ? ORDER BY id DESC LIMIT 1", user.ID).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
