package api

import (
	"gin_websocket/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

func UserRegister(c *gin.Context)  {
var userRegisterService service.UserRegisterService
if err :=c.ShouldBind(&userRegisterService);err!=nil{
	res:=userRegisterService.Register()
	c.JSON(http.StatusOK,res)
}else {
	c.JSON(http.StatusBadRequest,ErrorResponse(err))
	logging.Info(err)
}
}