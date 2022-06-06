package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Common.Response
	CommentList []Common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Common.Response
	Comment Common.Comment `json:"comment,omitempty"`
}

// CommentAction 评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_Type := c.Query("action_type")

	videoId, _ := strconv.ParseInt(video_id, 10, 64)
	actionType, _ := strconv.ParseInt(action_Type, 10, 64)

	var userid int64
	userid = 0
	//校验token是否有效
	if Common.CheckToken(token) {
		//从token中取出用户id
		userClaims, err := Common.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		userid = userClaims.ID
	} else {
		c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't login"})
	}

	if actionType == 1 {
		commentText := c.Query("comment_text")

		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Common.Response{StatusCode: 0},
			Comment:  service.CommentAction(videoId, userid, actionType, commentText, 0),
		})
	} else if actionType == 2 {
		comment_id := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(comment_id, 10, 64)

		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Common.Response{StatusCode: 0},
			Comment:  service.CommentAction(videoId, userid, actionType, "", commentId),
		})
	}

}

// CommentList 获取某视频的评论列表
func CommentList(c *gin.Context) {
	video_id := c.Query("video_id")
	token := c.Query("token")

	videoId, _ := strconv.ParseInt(video_id, 10, 64)

	var userid int64
	userid = 0
	//校验token是否有效
	if Common.CheckToken(token) {
		//从token中取出用户id
		userClaims, err := Common.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		userid = userClaims.ID
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Common.Response{StatusCode: 0},
		CommentList: service.CommentList(videoId, userid),
	})
}
