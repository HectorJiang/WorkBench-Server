package model

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ArticleId           uint         `gorm:"column:article_id;primary_key" json:"article_id"`
	ArticleTitle        string       `gorm:"column:article_title" json:"article_title"`
	ArticleContent      string       `gorm:"column:article_content;type:longtext" json:"article_content"`
	CategoryId          uint8        `gorm:"column:category_id" json:"category_id"`
	ArticlePic          string       `gorm:"column:article_pic" json:"article_pic"`
	ArticleView         int32        `gorm:"column:article_view" json:"article_view"`
	ArticleCommentCount int32        `gorm:"column:article_comment_count" json:"article_comment_count"`
	ArticleLikeCount    int32        `gorm:"column:article_like_count" json:"article_like_count"`
	ArticleStatus       int8         `gorm:"column:article_status" json:"article_status"`
	ArticlePrivilege    int8         `gorm:"column:article_privilege" json:"article_privilege"`
	ArticlePassword     string       `gorm:"column:article_password" json:"article_password"`
	CreatedAt           time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time    `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           sql.NullTime `gorm:"column:deleted_at" json:"deleted_at"`
}

type ArticleResult struct {
	Article
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

//Create article.
func CreateArticle(data *Article) int {
	fmt.Println(data)
	err := db.Create(&data).Error
	if err != nil {
		return 500
	}
	return 200
}

//Delete article
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("article_id=?", id).Delete(&article).Error
	if err != nil {
		return 500
	}
	return 200
}

//Edit article
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["article_title"] = data.ArticleTitle
	maps["article_content"] = data.ArticleContent
	maps["category_id"] = data.CategoryId
	maps["article_pic"] = data.ArticlePic
	maps["article_status"] = data.ArticleStatus
	maps["article_privilege"] = data.ArticlePrivilege
	maps["article_password"] = data.ArticlePrivilege
	err = db.Model(&article).Where("article_id=?", id).Updates(&maps).Error
	if err != nil {
		return 500
	}
	return 200
}

//Get one article
func GetArticle(id int) (ArticleResult, int) {
	var article Article
	var result ArticleResult
	err := db.Raw("select article.*,category.category_name from article, category where article.category_id = category.category_id and article_id = ?", id).Scan(&result).Error
	db.Model(&article).Where("article_id = ?", id).UpdateColumn("article_view", gorm.Expr("article_view + ?", 1))
	if err != nil {
		return result, 500
	}
	return result, 200
}

//Get article list
func GetArticleList() ([]ArticleResult, int, int64) {
	var articleList []ArticleResult
	var total int64
	total = 150
	err := db.Raw("select article.*,category.category_name from article,category where article.category_id = category.category_id").Scan(&articleList).Error
	if err != nil {
		return nil, 500, 0
	}
	return articleList, 200, total
}

//Get article list by category
// func GetArticleByCat(id int, pageSize int, pageNum int) ([]Article, int, int64) {
// 	var articleByCategory []ArticleResult
// 	var total int64

// 	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize)Find(&articleByCategory).Error

// 	db.Model(&articleByCategory).Where(
// 		"cid =?", id).Count(&total)
// 	if err != nil {
// 		return nil, 500, 0
// 	}
// 	return articleByCategory, 200, total
// }

//Get archive date

//Get article by archive date
