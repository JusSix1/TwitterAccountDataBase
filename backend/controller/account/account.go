package controller

import (
	"net/http"

	"github.com/JusSix1/TwitterAccountDataBase/entity"
	"github.com/gin-gonic/gin"
)

// GET /account/:email
func GetAccount(c *gin.Context) {
	var user entity.User
	var account []entity.Account

	email := c.Param("email")

	if tx := entity.DB().Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := entity.DB().Preload("Account_Status").Raw("SELECT * FROM accounts WHERE user_id = ?", user.ID).Find(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}
