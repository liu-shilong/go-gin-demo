package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-shilong/go-gin-demo/middleware/jwt"
	"github.com/liu-shilong/go-gin-demo/pkg/setting"
	"github.com/liu-shilong/go-gin-demo/routers/api"
	"github.com/liu-shilong/go-gin-demo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 标签
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 文章
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
