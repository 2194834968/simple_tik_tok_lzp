package main

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	Common.MysqlDb = Common.InitMysqlConnection()
	//test
	//service.Feed("1654003238325", 0)
	//service.Login("DemoUser", "123456")
	//service.Register("xiaoHua", "123456")
	//service.List(2)
	//Common.SaveCover("./public/15_陈奕迅-十面埋伏.mp4", "15_陈奕迅-十面埋伏.png", "./public/")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
