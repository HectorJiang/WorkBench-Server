package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hectorjiang/workbench/model"
)

//Create Category
func CreateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	//chang status into 0(nomal)
	// data.CategoryStatus = 1
	model.CreateCategory(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    data,
		"message": 200,
	})
}

//Delete Category
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": 200,
	})
}

//Edit Category
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	model.EditCategory(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": 200,
	})
}

//Get one Category
func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": code,
		"data":    data,
	})
}

//Get Category list
func GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCategoryList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": 200,
		"data":    data,
		"total":   total,
	})
}
