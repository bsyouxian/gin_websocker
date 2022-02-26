package service

import (
	"gin_websocket/model"
	"gin_websocket/serializer"
)

type UserRegisterService struct {
	UserName string `json:"user_name"form:"user_name"`
	Passwrod string `json:"passwrod" form:"password"`
}

//用户注册，返回一个序列化的值
func (service *UserRegisterService) Register() serializer.Response {
	var user model.User
	code := 200
	count := 0
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).Count(&count)
	if count!=0{
		return serializer.Response{
			Status: 400,
			Msg: "已存在，请更换。",
		}
	}
	user =model.User{
		UserName: service.UserName,
	}
	//加密密码
	if err := user.SetPassword(service.Passwrod);err!=nil{
		return serializer.Response{
			Status: 500,
			Msg: "加密出错",
		}
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status: code,
		Msg: "创建成功",
	}
}
