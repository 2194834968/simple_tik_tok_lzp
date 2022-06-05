package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	//校验token是否过期
	if !CheckToken(token) {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "Token expired"})
		return
	}

	//从token中取出用户id
	userClaims, err := ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if service.FavoriteAction(video_id, action_type, userClaims.ID) {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "action fail"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	if CheckToken(token) {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Common.Response{StatusCode: 0},
			VideoList: service.FavoriteList(userid),
		})
		return
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Common.Response{StatusCode: 1, StatusMsg: "登陆后才可查看喜欢的视频列表哦"},
			VideoList: nil,
		})
		return
	}
}
