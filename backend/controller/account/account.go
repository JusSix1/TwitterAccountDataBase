package controller

import (
	"net/http"

	"github.com/JusSix1/TwitterAccountDataBase/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /account
func CreateAccount(c *gin.Context) {
	var user entity.User
	var accountStatus entity.Account_Status
	var account entity.Account

	type AccountImport struct {
		Email_User        string
		Twitter_Account   string
		Twitter_Password  string
		Email_Accont      string
		Email_Password    string
		Phone_Number      string
		Years             uint `gorm:"required~Year is blank"`
		Account_Status_ID uint
	}

	var accountImport AccountImport

	if err := c.ShouldBindJSON(&accountImport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("email = ?", accountImport.Email_User).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", accountImport.Account_Status_ID).First(&accountStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account status not found"})
		return
	}

	if _, err := govalidator.ValidateStruct(accountImport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("user_id = ?", user.ID).Last(&account); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Last account not found"})
		return
	}

	// create new object for create new record
	newAccount := entity.Account{
		User_ID:           &user.ID,
		ID_Account:        account.ID_Account + 1,
		Twitter_Account:   accountImport.Twitter_Account,
		Twitter_Password:  accountImport.Twitter_Password,
		Email:             accountImport.Email_Accont,
		Email_Password:    accountImport.Email_Password,
		Phone_Number:      accountImport.Phone_Number,
		Years:             accountImport.Years,
		Account_Status_ID: &accountStatus.ID,
	}

	if err := entity.DB().Create(&newAccount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accountImport})

}

// GET /account/:email
func GetAccount(c *gin.Context) {
	var user entity.User
	var account []entity.Account

	email := c.Param("email")

	if tx := entity.DB().Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := entity.DB().Preload("Account_Status").Raw("SELECT * FROM accounts WHERE user_id = ? ORDER BY id DESC", user.ID).Find(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}
