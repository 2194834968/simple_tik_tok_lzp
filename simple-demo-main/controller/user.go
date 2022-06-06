package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	token, err := Common.GenerateToken(userid)
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
	token, err := Common.GenerateToken(userid)
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

	var myUserId int64
	myUserId = 0
	if Common.CheckToken(token) {
		//从token中取出用户id
		userClaims, err := Common.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			return
		}
		myUserId = userClaims.ID
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Common.Response{StatusCode: 0},
		User:     service.UserInfo(userid, myUserId),
	})
}
