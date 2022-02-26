package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strconv"
)

var(
	RedisClient *redis.Client
	RedisDB string
	RedisAddr string
	RedisPw string
	RedisDbName string
)

func Init()  {
	file, err := ini.Load("./conf/config.ini")//加载配置信息
	if err !=nil{
		fmt.Println("ini load failed",err)
	}
	LoadRedis(file)//读取配置信息
	Redis()
}
//Redis 在中间件中初始化redis链接
func Redis()  {
	db,_:=strconv.ParseUint(RedisDbName,10,64)
	Client:=redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB: int(db),
	})
	_,err := Client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
}
func LoadRedis(file *ini.File)  {
	RedisDB = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}