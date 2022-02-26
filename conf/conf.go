package conf

import (
	"context"
	"fmt"
	"gin_websocket/model"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	MongoDBClient *mongo.Client
	AppMode       string
	HttpPort      string
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string

	MongoDBName string
	MongoDBAddr string
	MongoDBPwd  string
	MongoDBPort string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err !=nil{
		fmt.Println("ini load failed",err)
	}
	loadServer(file)
	loadMysql(file)
	loadMongoDb(file)
	//mysql
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	MongoDB()//mongo
}
func MongoDB()  {
	//连接mongo
	clientOption:=options.Client().ApplyURI("mongodb://"+MongoDBAddr+":"+MongoDBPort)
	var err error
	MongoDBClient,err =mongo.Connect(context.TODO(),clientOption)
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	logging.Info("MongoDB connen successfully")

}
func loadServer(file *ini.File)  {
	AppMode=file.Section("service").Key("Appmode").String()
	HttpPort=file.Section("service").Key("HttpPort").String()

}
func loadMysql(file *ini.File)  {
	Db=file.Section("Mysql").Key("Db").String()
	DbHost=file.Section("Mysql").Key("DbHost").String()
	DbPort=file.Section("Mysql").Key("DbPort").String()
	DbUser=file.Section("Mysql").Key("DbUser").String()
	DbPassWord=file.Section("Mysql").Key("DbPassWord").String()
	DbName=file.Section("Mysql").Key("DbName").String()
}
func loadMongoDb(file *ini.File)  {
	MongoDBName=file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr=file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd=file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort=file.Section("MongoDB").Key("MongoDBPort").String()
}