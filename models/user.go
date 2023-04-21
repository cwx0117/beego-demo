package models

type User struct {
	Org   Org `gorm:"foreignKey:OrgId"`
	OrgId uint
	Id    int    `gorm:"primaryKey"`
	Name  string `gorm:"size(30)"`
	Pwd   string `gorm:"size(16)"`
	Email string `gorm:"uniqueIndex:idx_email_len_50,length:50"`
	Phone string `gorm:"uniqueIndex:idx_phone_len_11,length:11"`
}
