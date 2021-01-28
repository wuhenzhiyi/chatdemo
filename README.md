## 简介
	这是一个简单聊天室的小项目，使用Go+Gin+WebSocket+ElementUI技术，适合新手学习。
	目前并没有包含太多功能，只有简单的聊天功能，但这并不是终点，后续会陆续使用go相关的技术增加功能。
	我希望每个大版本是一个简单的版本，这样可以让大家和我一样渐进式的完成这个小项目。

## 演示地址
  [演示地址](http://demo.wuhen.site)

## 功能列表
- 互动聊天
- 用户列表

## 目录介绍

- app     -------------------主要逻辑文件
- router  -------------------路由文件
- static  -------------------静态文件包括css、js、iamge
- main.go -------------------启动文件


## 快速开始
	1、下载项目
		git clone https://github.com/wuhenzhiyi/chatdemo

	2、将项目放到自己的src目录下

	3、打开项目，修改地址为自己的ip地址和端口号，目前在template/index.html的connect方法修改，后面的版本会放到配置文件中。

	4、启动项目

	5、浏览器运行
	
## 依赖库下载
```go
go get 	github.com/gin-gonic/gin
go get  github.com/gorilla/websocket
```  
  
## 主要逻辑
### 路由实现 router/router.go
```go
	func SetRouter(){
		//新建服务
		router := gin.Default()

		//设置模板分隔符,以免和elemtnUi分隔符冲突
		//这个设置要在加载模板设置前面
		router.Delims("{{{", "}}}")

		//设置静态目录
		router.Static("/static", "static")

		//设置模板位置
		router.LoadHTMLGlob("template/*")

		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})

		//不写默认8080
		router.Run(":9999") 
	}
```

### 聊天界面设计  template/index.html
  目前界面不太美观，后续会美化
 

### 简单聊天互动  app/app.go
第一步 后端websocket代码
```go
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
    
```

第二步  前端连接websocket
```javascript
     connect(){
            let host = "192.168.30.138:9999"  //这里填写你自己的ip地址和端口
            let wsServer = 'ws://'+host+'/sendMessage'
            ws = new WebSocket(wsServer);

            ws.onopen = function(evt) {
                console.log("OPEN");
            }

            ws.onmessage = function(evt) {
              consoel.log(evt.data)
            };

            ws.onerror = function(evt) {
                console.log("ERROR: " + evt.data);
            };
        },
```

第三步  后端监听websocket连接
```go
    for {
			//读取消息
			_, recvMessage, err := conn.ReadMessage()

			if err != nil {
				//删除离开的用户
				fmt.Println(err)
				break
			}
			conn.WriteMessage(websocket.TextMessage, recvMessage)
		}
```

### 广播消息  app/app.go
以上只能把自己的消息发给自己，如果想把自己发的消息发给别人怎么办，就需要广播消息，将所有的用户连接存起来，然后发消息的时候，循环遍历所有的用户，给他们发消息。
```go
		type Client struct {
			Conn     *websocket.Conn
			UserName string
			Uid      string
		}

		var ClientMap map[string]Client

		//广播方法
		func broadcast(messageData MessageData) {
			jsonMessage, _ := json.Marshal(messageData)
			for _, c := range ClientMap {
				c.Conn.WriteMessage(websocket.TextMessage, jsonMessage)
			}
		}
```

### 待完善空间
- 文档写得不好还请见谅，大家多多提建议
- 聊天室功能：后续可以加入聊天室的选择，可以切换聊天室聊天
- 注册/登录功能，聊天记录功能，可以练习gorm+mysql+redis
- 或者更有趣点的考虑增加机器人聊天等等

### 友情推荐
[Go小考](https://www.golangroadmap.com/)