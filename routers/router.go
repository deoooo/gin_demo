package routers

import (
	"github.com/deoooo/gin_demo/middleware/jwt"
	"github.com/deoooo/gin_demo/pkg/setting"
	v1 "github.com/deoooo/gin_demo/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		// 登陆接口
		apiV1.GET("/login", v1.Login)
	}

	v1User := apiV1.Group("/user")
	v1User.Use(jwt.JWT())
	{
		v1User.GET("/hello", v1.Hello)
	}

	return r
}
