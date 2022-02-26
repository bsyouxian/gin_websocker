package service

import (
	"context"
	"fmt"
	"gin_websocket/conf"
	"gin_websocket/model/ws"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	//插入到mongDB中
	collection := conf.MongoDBClient.Database(database).Collection(id) //如果没id这个集合会创建一个
	conmmeet := ws.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), conmmeet)
	return err
}
func FindMany(database, sendId, id string, time int64, pagaSize int) (resule []ws.Result, err error) {
	var resuleMe []ws.Trainer
	var resuletYou []ws.Trainer
	sendIDCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	sendIDTimecurcor, err := sendIDCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pagaSize)))
	idTimecurcor, err := idCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pagaSize)))
	err = sendIDTimecurcor.All(context.TODO(), &resuletYou) // sendId 对面发过来的
	err = idTimecurcor.All(context.TODO(), &resuleMe)       // Id 发给对面的
	resule, _ = AppendAndSort(resuleMe, resuletYou)
	return
}
func AppendAndSort(resuleMe, resuletYou []ws.Trainer) (results []ws.Result, err error) {
	for _, r := range resuleMe {
		sendSort := SendSortMsg{ //构造返回的msg
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		resule := ws.Result{ //走早饭会的内容，包括传送这
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, resule)
	}
	for _, r := range resuletYou {
		sendSort := SendSortMsg{ //构造返回的msg
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		resule := ws.Result{ //走早饭会的内容，包括传送这
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "YOU",
		}
		results = append(results, resule)
	}
	return
}
