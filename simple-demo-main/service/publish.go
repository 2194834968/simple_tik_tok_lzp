package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Common"
	"strings"
	"time"
)

//将视频记录提交到数据库
func Publish(AuthorId int64, finalName string) bool {
	db := Common.MysqlConnection()

	videoFileaddress := "http://10.30.23.55:8080/static/"
	PlayUrl := fmt.Sprintf("%s%s.mp4", videoFileaddress, finalName)
	CoverUrl := fmt.Sprintf("%s%s.png", videoFileaddress, finalName)

	//检查视频是否已经投稿
	var PlayUrlExist string
	db.Table("videos").Select("play_url").Where("play_url = ?", PlayUrl).Take(&PlayUrlExist)
	if strings.EqualFold(PlayUrl, PlayUrlExist) {
		return false
	}

	videoDatebase := VideoDatebase{
		Title:          finalName,
		Author_Id:      AuthorId,
		Play_Url:       PlayUrl,
		Cover_Url:      CoverUrl,
		Favorite_Count: 0,
		Comment_Count:  0,
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}

	db.Table("videos").Create(&videoDatebase)

	//检查视频投稿是否成功
	var PlayUrlTest string
	db.Table("videos").Select("play_url").Where("play_url = ?", PlayUrl).Take(&PlayUrlTest)
	if PlayUrlTest == "" {
		return false
	}
	return true
}

func List(userId int64) []Common.Video {
	db := Common.MysqlConnection()

	var VideoList []Common.Video
	//用于返回的视频列表

	var VideoListTemp []VideoDatebase
	//用于数据库查询的视频列表

	db.Table("videos").Where("Author_Id = ?", userId).Find(&VideoListTemp)

	for i := 0; i < len(VideoListTemp); i++ {

		var videoTemp Common.Video

		videoTemp.Id = VideoListTemp[i].Id
		videoTemp.Title = VideoListTemp[i].Title
		videoTemp.Author = UserInfo(VideoListTemp[i].Author_Id)
		videoTemp.PlayUrl = VideoListTemp[i].Play_Url
		videoTemp.CoverUrl = VideoListTemp[i].Cover_Url
		videoTemp.FavoriteCount = VideoListTemp[i].Favorite_Count
		videoTemp.CommentCount = VideoListTemp[i].Comment_Count

		var FavoriteId int64
		db.Table("favorite").Select("Id").Where("video_id = ? AND user_id = ?", videoTemp.Id, userId).Find(&FavoriteId)
		if FavoriteId != 0 {
			videoTemp.IsFavorite = true
		} else {
			videoTemp.IsFavorite = false
		}

		VideoList = append(VideoList, videoTemp)
		//fmt.Printf("\n playAuthor %s \n", videoTemp.Author.Name)
	}
	return VideoList
}
