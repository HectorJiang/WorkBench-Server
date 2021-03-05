package model

type Myself struct {
	MyselfDescription string `gorm:"column:myself_description" json:"myself_description"`
}
