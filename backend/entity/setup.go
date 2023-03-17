package entity

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("Enteracc.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		// User
		&Gender{},
		&User{},

		//Account
		&Account_Status{},
		&Account{},
	)

	db = database

	Gen1 := Gender{
		Gender: "Male",
	}
	Gen2 := Gender{
		Gender: "Female",
	}
	Gen3 := Gender{
		Gender: "Other",
	}
	db.Model(&Gender{}).Create(&Gen1)
	db.Model(&Gender{}).Create(&Gen2)
	db.Model(&Gender{}).Create(&Gen3)

	User1 := User{
		Email:           "natt@gmail.com",
		Password:        "$2a$12$I0y6Rso/myQzK0EXsS0dv.a908//LMR7faAJgUJ.7LY2GrzoEsvWa",
		FirstName:       "นัท",
		LastName:        "ทัน",
		Profile_Name:    "Udong",
		Phone_number:    "0123456789",
		Birthday:        time.Now(),
		Profile_Picture: "data:image/jpeg;base64,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",
		Gender_ID:       &Gen1.ID,
	}
	db.Model(&User{}).Create(&User1)
	User2 := User{
		Email:           "AB@gmail.com",
		Password:        "$2a$12$I0y6Rso/myQzK0EXsS0dv.a908//LMR7faAJgUJ.7LY2GrzoEsvWa",
		FirstName:       "พชรพล",
		LastName:        "วรสุนทร",
		Profile_Name:    "๋JusSix",
		Phone_number:    "0123456789",
		Birthday:        time.Now(),
		Profile_Picture: "data:image/jpeg;base64,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",
		Gender_ID:       &Gen1.ID,
	}
	db.Model(&User{}).Create(&User2)

	AS1 := Account_Status{
		Status: "Sold",
	}
	AS2 := Account_Status{
		Status: "Unsold",
	}
	db.Model(&Account_Status{}).Create(&AS1)
	db.Model(&Account_Status{}).Create(&AS2)

	Account1 := Account{
		ID_Account:        1,
		User_ID:           &User1.ID,
		Twitter_Account:   "Account_Name_1",
		Twitter_Password:  "Account_Password_1",
		Email:             "Email_Name_1",
		Email_Password:    "Email_Password_1",
		Phone_Number:      "Phone_1",
		Years:             2018,
		Account_Status_ID: &AS1.ID,
	}
	Account2 := Account{
		ID_Account:        2,
		User_ID:           &User1.ID,
		Twitter_Account:   "Account_Name_2",
		Twitter_Password:  "Account_Password_2",
		Email:             "Email_Name_2",
		Email_Password:    "Email_Password_2",
		Phone_Number:      "Phone_2",
		Years:             2019,
		Account_Status_ID: &AS2.ID,
	}
	Account3 := Account{
		ID_Account:        3,
		User_ID:           &User1.ID,
		Twitter_Account:   "Account_Name_3",
		Twitter_Password:  "Account_Password_3",
		Email:             "Email_Name_3",
		Email_Password:    "Email_Password_3",
		Phone_Number:      "Phone_3",
		Years:             2018,
		Account_Status_ID: &AS2.ID,
	}
	Account4 := Account{
		ID_Account:        4,
		User_ID:           &User1.ID,
		Twitter_Account:   "Account_Name_4",
		Twitter_Password:  "Account_Password_4",
		Email:             "Email_Name_4",
		Email_Password:    "Email_Password_4",
		Phone_Number:      "Phone_4",
		Years:             2020,
		Account_Status_ID: &AS2.ID,
	}
	db.Model(&Account{}).Create(&Account1)
	db.Model(&Account{}).Create(&Account2)
	db.Model(&Account{}).Create(&Account3)
	db.Model(&Account{}).Create(&Account4)
}
