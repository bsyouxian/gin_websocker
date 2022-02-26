package service

import (
	"encoding/json"
	"fmt"
	"gin_websocket/conf"
	"gin_websocket/pkg/e"
	"github.com/gorilla/websocket"
)

func (massge *ClientManager) Start() {
	for {
		fmt.Println("——————————管道通信————————")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新连接，%v", conn.ID)
			Manager.Clients[conn.ID] = conn
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Printf("连接失败：%v", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast: //查看对方是否在线
			massge := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false                           //默认不在线
			for id, conn := range Manager.Clients { //去客户端里找对方的id，
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- massge:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
				id := broadcast.Client.ID //发送方的id
				if flag {
					replyMsg := &ReplyMsg{
						Code:    e.WebsocketOnlineReply,
						Content: "对方已应答",
					}
					msg, _ := json.Marshal(replyMsg)
					_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
					err := InsertMsg(conf.MongoDBName, id, string(massge), 1, int64((3 * month))) //1已经阅读
					if err != nil {
						fmt.Println("InsetOne err", err)
					}
				}
			}
		}

	}

}
