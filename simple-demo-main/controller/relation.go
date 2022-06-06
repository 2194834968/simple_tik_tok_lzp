package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Common.Response
	UserList []Common.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	to_user_id := c.Query("to_user_id")
	action_type := c.Query("action_type")

	toUserId, err := strconv.ParseInt(to_user_id, 10, 64)
	actionType, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Common.Response{StatusCode: 2, StatusMsg: "The user_id format is incorrect,user_id must be pure numbers"},
		})
		panic("The user_id format is incorrect,user_id must be pure numbers")
	}

	if Common.CheckToken(token) {
		//从token中取出用户id
		userClaims, err := Common.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		if service.RelationAction(userClaims.ID, toUserId, actionType) {
			c.JSON(http.StatusOK, Common.Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "登陆后才能关注别人哦"})
	}
}

// FollowList 获取userId关注的人
func FollowList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	userId, err := strconv.ParseInt(user_id, 10, 64)
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
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		myUserId = userClaims.ID
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Common.Response{StatusCode: 0},
		UserList: service.RelationFollowList(userId, myUserId),
	})
}

// FollowerList 获取关注userId的人
func FollowerList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	userId, err := strconv.ParseInt(user_id, 10, 64)
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
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		myUserId = userClaims.ID
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Common.Response{StatusCode: 0},
		UserList: service.RelationFollowerList(userId, myUserId),
	})
}
