import { GendersInterface } from "./IGender"

export interface UsersInterface { // Maybe BUG warning คือตัวเชื่อม interface วนไปมา
    ID: number,
    Email: string,
	Password: string,
	Profile_Name: string,
	Profile_Picture: string, // JavaScript Int8Array = byte ; แต่ต้องลองใช้ดูว่า work ไหมกับ go []byte
	Birthday: Date,
	Gender_ID: number,
	Gender: GendersInterface,
}