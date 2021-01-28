package app

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	UserName string
	Uid      string
	UserHead string
}

type User struct {
	UserName string `json:"userName"`
	UserHead string `json:"userHead"`
}

var ClientMap map[string]Client

type MessageData struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type Message struct {
	Content  string `json:"content"`
	UserName string `json:"userName"`
	UserHead string `json:"userHead"`
}

func init() {
	ClientMap = make(map[string]Client)
}
