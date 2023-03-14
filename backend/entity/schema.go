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
	Email           string    `gorm:"uniqueIndex" valid:"email~รูปแบบ email ไม่ถูกต้อง,required~กรุณากรอก email"`
	Password        string    `valid:"minstringlength(8)~ความยาวรหัสผ่านต้องไม่ต่ำกว่า 8 ตัวอักษร,required~กรุณากรอกรหัสผ่าน"`
	Profile_Name    string    `valid:"maxstringlength(50)~ชื่อความยาวไม่เกิน 50 ตัวอักษร,required~กรุณากรอกชื่อ"`
	Profile_Picture string    `valid:"image_valid~รูปภาพไม่ถูกต้อง"`
	Birthday        time.Time `valid:"NotFutureTime~วันเกิดต้องไม่เป็นอนาคต,MoreThan18YearsAgo~คุณต้องมีอายุมากกว่า 18 ปี"`
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
