package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	TagId     uint16       `gorm:"column:tag_id" json:"tag_id"`
	TagName   string       `gorm:"column:tag_name" json:"tag_name"`
	CreatedAt time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at" json:"deleted_at"`
}

//Create tag
func CreateTag(data *Tag) int {
	err := db.Create(&data).Error
	if err != nil {
		return 500
	}
	return 200
}

//Delete tag
func DeleteTag(id int) int {
	var tag Tag
	err := db.Where("tag_id", id).Delete(&tag).Error
	if err != nil {
		return 500
	}
	return 200
}

//Edit tag
func EditTag(id int, data *Tag) int {
	var tag Tag
	var maps = make(map[string]interface{})
	maps["tag_name"] = data.TagName
	err = db.Model(&tag).Where("tag_id=?", id).Updates(maps).Error
	if err != nil {
		return 500
	}
	return 200
}

//Get tag list
func GetTagList(pagesize int, pagenum int) ([]Tag, int64) {
	var tag []Tag
	var total int64
	err = db.Find(&tag).Limit(pagesize).Offset((pagenum - 1) * pagesize).Error
	db.Model(&tag).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return tag, total
}
