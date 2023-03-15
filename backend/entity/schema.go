package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// ----------User----------
type Gender struct {
	gorm.Model
	Gender string
	User   []User `gorm:"foreignKey:Gender_ID"`
}

type User struct {
	gorm.Model
	Email           string    `gorm:"uniqueIndex" valid:"email~Invalid Email format,required~Email is blank"`
	FirstName       string    `valid:"required~First name is blank"`
	LastName        string    `valid:"required~Last name is blank"`
	Password        string    `valid:"minstringlength(8)~Password must be longer than 8 characters,required~Password is blank"`
	Profile_Name    string    `valid:"maxstringlength(50)~Must be no more than 50 characters long,required~Profile name is blank"`
	Profile_Picture string    `valid:"image_valid~Please change the picture"`
	Birthday        time.Time `valid:"NotFutureTime~The day must not be the future,MoreThan18YearsAgo~You must be over 18 years old"`
	Phone_number    string    `valid:"required~Phone number is blank,matches([0-9]{10})~Phone number invalid format"`
	Gender_ID       *uint     `valid:"-"`
	Gender          Gender    `gorm:"references:id" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("DelayNow10Min", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Minute * -10))
	})

	govalidator.CustomTypeTagMap.Set("NotFutureTime", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return !t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("MoreThan18YearsAgo", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		ageLimit := time.Now().AddDate(-18, 0, 0)
		return t.Before(ageLimit)
	})

	govalidator.TagMap["image_valid"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^(data:image(.+);base64,.+)$")
	})
}
