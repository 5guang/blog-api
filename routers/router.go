package routers

import (
	u "blog/api/user"
	v1 "blog/api/v1"
	_ "blog/docs"
	"blog/middleware/jwt"
	"blog/middleware/logger"
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", v1.Ping)
	// 中间件
	r.Use(logger.LogerMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// 路由
	//apiUser := r.Group("/user/")
	apiV1 := r.Group("/api/v1")
	user := r.Group("api/user")

	{
		//apiUser.POST("/login", api.Login)
		user.POST("/register", u.Register)
		user.POST("/login", u.Login)
	}


	{
		// 获取所有标签
		apiV1.GET("/tags", v1.GetTags)
		// 通过文章id获取tags
		apiV1.POST("/getTagsByArticleId", v1.GetTagsByArticleId)
		// 获取所有文章
		apiV1.POST("/getArticleList", v1.GetArticles)
		// 通过标签id获取文章
		apiV1.POST("/getArticlesByTagId", v1.GetArticlesByTagId)
		// 通过文章id获取文章
		apiV1.POST("/getArticleById", v1.GetArticleByArticleId)
	}
	apiV1.Use(jwt.Jwt())
	{
		// 添加文章
		apiV1.POST("/article", v1.AddArticle)
		// 更改文章
		apiV1.PUT("/article", v1.UpdateArticle)
		// 删除文章
		apiV1.DELETE("/article", v1.DelArticle)
		// 添加标签
		apiV1.POST("/tags", v1.AddTag)
		// 删除标签
		apiV1.DELETE("/tags", v1.DelTag)
		// 更改标签
		apiV1.PUT("/tags", v1.UpdateTag)
	}
	return r
}
