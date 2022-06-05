package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]Common.User{
	"zhangleidouyin": {
		Id:             1,
		Name:           "zhanglei",
		Follow_Count:   10,
		Follower_Count: 5,
		IsFollow:       true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Common.Response
	User Common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userid, state := service.Register(username, password)
	token, err := GenerateToken(userid)
	if state == false {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 2, StatusMsg: "false to GenerateToken"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Common.Response{StatusCode: 0},
			UserId:   userid,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userid, exist := service.Login(username, password)
	token, err := GenerateToken(userid)
	if exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Common.Response{StatusCode: 0},
			UserId:   userid,
			Token:    token,
		})
	} else if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 2, StatusMsg: "false to GenerateToken"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	userid, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 2, StatusMsg: "The user_id format is incorrect,user_id must be pure numbers"},
		})
		panic("The user_id format is incorrect,user_id must be pure numbers")
	}

	if CheckToken(token) {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 0},
			User:     service.UserInfo(userid),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}