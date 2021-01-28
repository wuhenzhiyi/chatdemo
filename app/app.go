package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func SendMessageApi(ctx *gin.Context) {

	//初始化websocket对象
	upGrader := websocket.Upgrader{
		//跨域设置
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//建立连接
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("已建立连接")

	//关闭连接
	defer conn.Close()

	//记录用户信息

	//使用时间戳当用户id
	userTimeStamp := int(time.Now().UnixNano())
	uid := strconv.Itoa(userTimeStamp)

	//随机获取用户名
	userName := getRandUserName()

	//随即获取用户头像
	userHeadIndex := rand.Intn(9) + 1
	userHead := strconv.Itoa(userHeadIndex) + ".jpg"

	client := Client{
		Conn:     conn,
		UserName: userName,
		Uid:      uid,
		UserHead: userHead,
	}

	//将用户加入用户组,使用时间戳当用户的键值
	ClientMap[client.Uid] = client

	//新用户加入发送欢迎消息
	//封装消息体
	dataMap := make(map[string]interface{})
	dataMap["content"] = "欢迎 " + client.UserName + " 加入聊天室"
	dataMap["userName"] = "system"
	dataMap["userList"] = getUserList()

	messageData := MessageData{
		Action: "systemMessage",
		Data:   dataMap,
	}

	broadcast(messageData)

	//监听消息
	for {
		//读取消息
		_, recvMessage, err := conn.ReadMessage()

		if err != nil {
			//删除离开的用户
			fmt.Println(err)
			delete(ClientMap, client.Uid)

			//封装消息体
			dataMap := make(map[string]interface{})
			dataMap["content"] = client.UserName + " 离开了聊天室"
			dataMap["userName"] = "system"
			dataMap["userList"] = getUserList()

			messageData := MessageData{
				Action: "systemMessage",
				Data:   dataMap,
			}

			broadcast(messageData)
			break
		}

		//封装消息体
		message := Message{
			Content:  string(recvMessage),
			UserName: client.UserName,
			UserHead: client.UserHead,
		}

		messageData := MessageData{
			Action: "userMessage",
			Data:   message,
		}

		//广播消息
		broadcast(messageData)
	}
}

//广播
func broadcast(messageData MessageData) {
	jsonMessage, _ := json.Marshal(messageData)
	for _, c := range ClientMap {
		c.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
	}
}

//随机获取用户名字
func getRandUserName() string {
	//形容词
	adjectiveList := []string{
		"开心的", "高兴的", "愤怒的", "愉快的", "伤心的", "憨憨的", "聪明的",
	}

	nounList := []string{
		"兔子", "老虎", "青蛙", "小鸟", "蚂蚁", "大象", "狮子", "小猪", "小猫", "小狗", "小马驹",
	}

	adjectiveIndex := rand.Intn(len(adjectiveList))
	adjective := adjectiveList[adjectiveIndex]

	nounIndex := rand.Intn(len(nounList))
	noun := nounList[nounIndex]

	userName := adjective + noun
	return userName
}

func getUserList() []User {
	var userList []User

	for _, c := range ClientMap {
		user := User{
			UserName: c.UserName,
			UserHead: c.UserHead,
		}
		userList = append(userList, user)
	}

	return userList
}
