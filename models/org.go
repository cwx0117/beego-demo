package models

type Org struct {
	Id       int64  `gorm:"primaryKey"`
	Province string `gorm:"size(18)"`
	City     string `gorm:"size(18)"`
	County   string `gorm:"size(18)"`
	AreaCode string `gorm:"uniqueIndex:idx_area_code_len_50,length:50"`
}
