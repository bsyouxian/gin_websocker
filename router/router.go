package router

import (
	"gin_websocket/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_websocket/service"
)

func NewRouter() *gin.Engine {
	r:=gin.Default()
	r.Use(gin.Recovery(),gin.Logger())
//Recovery恢复用的，looger是日志
v1:=r.Group("/")
{
	v1.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,"success")
	})
	v1.POST("user/register",api.UserRegister)
	v1.GET("ws",service.Handller)
}
return r
}