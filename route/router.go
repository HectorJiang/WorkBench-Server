package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hectorjiang/workbench/api/v1"
	"github.com/hectorjiang/workbench/middleware"
	"github.com/hectorjiang/workbench/util"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(middleware.Cors())
	router := r.Group("api/v1")
	{
		//Article route
		//For front-end
		router.GET("article/detail/:id", v1.GetArticle)
		router.GET("article", v1.GetArticleList)
		router.GET("article/list/:id", v1.GetArticleByCat)

		//For back-end(no jwt at present)
		router.POST("admin/article/add", v1.CreateArticle)
		router.DELETE("admin/article/:id", v1.DeleteArticle)
		router.PUT("admin/article/:id", v1.EditArticle)
		router.GET("admin/article/detail/:id", v1.GetArticle)
		router.GET("admin/article", v1.GetArticleList)

		//Article category
		//For front-end
		router.GET("category/:id", v1.GetCategory)
		router.GET("category", v1.GetCategoryList)

		//For back-end
		router.POST("admin/category/add", v1.CreateCategory)
		router.DELETE("admin/category/:id", v1.DeleteCategory)
		router.PUT("admin/category/:id", v1.EditCategory)
		router.GET("admin/category", v1.GetCategoryList)

	}

	_ = r.Run(util.HttpPort)
}
