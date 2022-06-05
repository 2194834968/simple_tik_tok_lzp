package controller

import (
	"github.com/RaymondCode/simple-demo/Common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	Common.Response
	VideoList []Common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")

	var userid int64
	userid = 0
	//校验token是否有效
	if CheckToken(token) {
		//从token中取出用户id
		userClaims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
			return
		}
		userid = userClaims.ID
	}

	//fmt.Printf("\n %s \n", latestTime)
	serviceResult := service.Feed(latestTime, userid)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Common.Response{StatusCode: 0},
		VideoList: serviceResult.VideoList,
		NextTime:  serviceResult.NextTime,
	})
}
