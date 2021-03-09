package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hectorjiang/workbench/model"
)

//Create article
func CreateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": 200,
	})
}

//Delete article
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": 200,
	})
}

//Edit article
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": 200,
	})
}

//Get one article
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": 200,
	})
}

//Get article list
func GetArticleList(c *gin.Context) {
	// pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	// pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// title := c.Query("title")

	// switch {
	// case pageSize >= 100:
	// 	pageSize = 100
	// case pageSize <= 0:
	// 	pageSize = 10
	// }

	// if pageNum == 0 {
	// 	pageNum = 1
	// }

	data, code, total := model.GetArticleList()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": 200,
	})
}

//Get article list by category
func GetArticleByCat(c *gin.Context) {

}
