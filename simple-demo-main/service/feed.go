package service

import (
	"github.com/RaymondCode/simple-demo/Common"
	"strconv"
	"time"
)

type FeedServiceResponse struct {
	VideoList []Common.Video
	NextTime  int64
}

//用于数据库查询的视频列表结构体
type VideoDatebase struct {
	Id             int64
	Title          string
	Author_Id      int64
	Play_Url       string
	Cover_Url      string
	Favorite_Count int64
	Comment_Count  int64
	Created_At     time.Time
	Updated_At     time.Time
}

func Feed(latestTimeString string, userId int64) FeedServiceResponse {

	db := Common.MysqlDb
	//fmt.Printf("什么大苏打大苏打大啊撒打算大苏打大苏打啊大大实打实的")
	latest_time := time.Unix(0, 0)
	//转换时间戳
	if latestTimeString != "" {
		latestTimeInt, _ := strconv.ParseInt(latestTimeString, 10, 64)
		latest_time = time.Unix(0, latestTimeInt)
		//fmt.Printf("\n %s \n", latest_time.String())
	}

	var VideoList []Common.Video
	//用于返回的视频列表

	var VideoListTemp []VideoDatebase
	//用于数据库查询的视频列表

	var NextTime time.Time
	//最早时间戳

	db.Table("videos").Where("created_at > ?", latest_time).Find(&VideoListTemp)

	for i := 0; i < 30 && i < len(VideoListTemp); i++ {

		var videoTemp Common.Video

		if i == 0 {
			NextTime = VideoListTemp[i].Created_At
		}
		videoTemp.Id = VideoListTemp[i].Id
		videoTemp.Title = VideoListTemp[i].Title
		videoTemp.Author = UserInfo(VideoListTemp[i].Author_Id, userId)
		videoTemp.PlayUrl = VideoListTemp[i].Play_Url
		videoTemp.CoverUrl = VideoListTemp[i].Cover_Url
		videoTemp.FavoriteCount = VideoListTemp[i].Favorite_Count
		videoTemp.CommentCount = VideoListTemp[i].Comment_Count

		if userId != 0 {
			var FavoriteId int64
			db.Table("favorite").Select("Id").Where("video_id = ? AND user_id = ?", videoTemp.Id, userId).Find(&FavoriteId)
			if FavoriteId != 0 {
				videoTemp.IsFavorite = true
			} else {
				videoTemp.IsFavorite = false
			}
		} else {
			videoTemp.IsFavorite = false
		}

		VideoList = append(VideoList, videoTemp)
	}

	return FeedServiceResponse{
		VideoList: VideoList,
		NextTime:  NextTime.Unix(),
	}
}
