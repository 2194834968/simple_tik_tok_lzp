package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	//校验token是否过期
	if !Common.CheckToken(token) {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "Token expired"})
		return
	}

	//从token中取出用户id
	userClaims, err := Common.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//数据格式转换
	videoId, _ := strconv.ParseInt(video_id, 10, 64)
	actionType, _ := strconv.ParseInt(action_type, 10, 64)

	if service.FavoriteAction(videoId, actionType, userClaims.ID) {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "action fail"})
	}
}

// FavoriteList 用于获取用户喜欢的视频ID列表
func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")

	userId, _ := strconv.ParseInt(userid, 10, 64)

	if Common.CheckToken(token) {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Common.Response{StatusCode: 0},
			VideoList: service.FavoriteList(userId),
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
