package test

import (
	"testing"

	"github.com/hectorjiang/workbench/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestDB(t *testing.T) {
	db, err := gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Auto migrate
	db.AutoMigrate(&model.Category{})

	c1 := model.Category{1, "category01"}
	c2 := model.Category{2, "category02"}

	//Create data
	db.Create(&c1)
	db.Create(&c2)
	t.Log("Test is over.")

}
