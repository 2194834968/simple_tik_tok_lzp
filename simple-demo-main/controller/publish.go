package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Common.Response
	VideoList []Common.Video `json:"video_list"`
}

// Publish 上传视频
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	fileTitle := c.PostForm("title")

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

	//取出视频数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	//建立保存路径
	finalName := fmt.Sprintf("%d_%s", userClaims.ID, fileTitle)
	finalNameMp4 := fmt.Sprintf("%d_%s.mp4", userClaims.ID, fileTitle)
	finalNamePng := fmt.Sprintf("%d_%s.png", userClaims.ID, fileTitle)

	//public\15_暗号-周杰伦
	//fmt.Printf(saveFile)

	//服务器保存视频数据
	saveFile := filepath.Join("./public/", finalNameMp4)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	//服务器保存视频封面
	if !Common.SaveCover(saveFile, finalNamePng, "./public/") {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	//数据库保存记录
	if !service.Publish(userClaims.ID, finalName) {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "video already exist"})
		return
	}

	c.JSON(http.StatusOK, Common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList 某用户的视频上传列表
func PublishList(c *gin.Context) {
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
			c.JSON(http.StatusOK, VideoListResponse{
				Response:  Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
				VideoList: nil,
			})
			return
		}
		myUserId = userClaims.ID
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Common.Response{StatusCode: 0},
		VideoList: service.List(userid, myUserId),
	})
}
