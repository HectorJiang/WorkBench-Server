package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

/**
Category model
*/

type Category struct {
	CategoryId     uint8        `gorm:"column:category_id;primary_key" json:"category_id"`
	CategoryName   string       `gorm:"column:category_name" json:"category_name"`
	CategoryStatus int8         `gorm:"column:category_status" json:"category_status"`
	CreatedAt      time.Time    `gorm:"column:created_at" json:"	"`
	UpdatedAt      time.Time    `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      sql.NullTime `gorm:"column:deleted_at" json:"deleted_at"`
}

//Check category
// func CheckCategory(name string) (code int) {
// 	var category Category
// 	db.Select("id").Where("name=?", name).First(&category)
// }

//Create category
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return 500
	}
	return 200
}

//Delete category
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("category_id=?", id).Delete(&category).Error
	if err != nil {
		return 500
	}
	return 200
}

//Edit category
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["category_name"] = data.CategoryName
	maps["category_status"] = data.CategoryStatus
	err = db.Model(&category).Where("category_id=?", id).Updates(maps).Error
	if err != nil {
		return 500
	}
	return 200
}

//Get one category
func GetCategory(id int) (Category, int) {
	var category Category
	db.Where("category_id=?", id).First(&category)
	return category, 200
}

//Get category list
func GetCategoryList(pagesize int, pagenum int) ([]Category, int64) {
	var category []Category
	var total int64
	err = db.Find(&category).Limit(pagesize).Offset((pagenum - 1) * pagesize).Error
	db.Model(&category).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return category, total
}
