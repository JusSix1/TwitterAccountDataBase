package controller

import (
	"net/http"

	"github.com/JusSix1/TwitterAccountDataBase/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /users --> ใช้อยู่
func CreateUser(c *gin.Context) {
	var user entity.User
	var emailCheck entity.User
	var gender entity.Gender

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", user.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาเลือกเพศ"})
		return
	}

	if tx := entity.DB().Where("email = ?", user.Email).First(&emailCheck); !(tx.RowsAffected == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "อีเมลนี้ถูกใช้ไปแล้ว"})
		return
	}

	// create new object for create new record
	newUser := entity.User{
		Email:           user.Email,
		Password:        user.Password,
		Profile_Name:    user.Profile_Name,
		Profile_Picture: user.Profile_Picture,
		Gender:          gender,
		Birthday:        user.Birthday,
	}

	// validate user
	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hashing after validate
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	newUser.Password = string(hashPassword)

	if err := entity.DB().Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}
