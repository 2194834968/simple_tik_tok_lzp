package service

import (
	"github.com/RaymondCode/simple-demo/Common"
)

//用于数据库查询的点赞结构体
type FavoriteDatebase struct {
	Id       int64
	Video_Id int64
	User_Id  int64
}

//用于获取UserId喜欢的视频列表
func FavoriteList(userId int64, myUserId int64) []Common.Video {
	db := Common.MysqlDb

	var VideoIdList []int64
	//用于获取用户喜欢的视频ID列表
	db.Table("favorite").Select("video_id").Where("user_id = ?", userId).Find(&VideoIdList)

	var VideoList []Common.Video
	//用于返回的视频列表

	for i := 0; i < len(VideoIdList); i++ {
		var videoDatebaseTemp VideoDatebase
		//用于数据库查询的视频结构

		var videoTemp Common.Video
		//用于返回的视频结构

		db.Table("videos").Where("Id = ?", VideoIdList[i]).Find(&videoDatebaseTemp)

		videoTemp.Id = videoDatebaseTemp.Id
		videoTemp.Title = videoDatebaseTemp.Title
		videoTemp.Author = UserInfo(videoDatebaseTemp.Author_Id, myUserId)
		videoTemp.PlayUrl = videoDatebaseTemp.Play_Url
		videoTemp.CoverUrl = videoDatebaseTemp.Cover_Url
		videoTemp.FavoriteCount = videoDatebaseTemp.Favorite_Count
		videoTemp.CommentCount = videoDatebaseTemp.Comment_Count

		if myUserId != 0 {
			var FavoriteId int64
			db.Table("favorite").Select("Id").Where("video_id = ? AND user_id = ?", videoTemp.Id, myUserId).Find(&FavoriteId)
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
	return VideoList
}

func FavoriteAction(video_id int64, action_type int64, userId int64) bool {
	db := Common.MysqlDb

	favoriteDatebase := FavoriteDatebase{
		Video_Id: video_id,
		User_Id:  userId,
	}

	//更新视频表里的视频点赞数
	var favorite_Count_old int64
	var Author_Id int64
	db.Table("videos").Select("Author_Id").Where("Id = ?", video_id).Take(&Author_Id)
	db.Table("videos").Select("Favorite_Count").Where("Id = ?", video_id).Take(&favorite_Count_old)

	//更新用户表里的点赞总数
	//视频作者
	var Total_Favorite_Users_old int64
	db.Table("users").Select("Total_Favorite").Where("Id = ?", Author_Id).Take(&Total_Favorite_Users_old)

	//点赞者
	var Favorite_Count_Users_old int64
	db.Table("users").Select("Favorite_Count").Where("Id = ?", userId).Take(&Favorite_Count_Users_old)

	if action_type == 1 {
		db.Table("favorite").Create(&favoriteDatebase)
		db.Table("videos").Where("Id = ?", video_id).Update("Favorite_Count", favorite_Count_old+1)
		db.Table("users").Where("Id = ?", Author_Id).Update("Total_Favorite", Total_Favorite_Users_old+1)
		db.Table("users").Where("Id = ?", userId).Update("Favorite_Count", Favorite_Count_Users_old+1)

	} else if action_type == 2 {
		db.Table("favorite").Where("video_id = ? AND user_id = ?", video_id, userId).Delete(&favoriteDatebase)
		db.Table("videos").Where("Id = ?", video_id).Update("Favorite_Count", favorite_Count_old-1)
		db.Table("users").Where("Id = ?", Author_Id).Update("Total_Favorite", Total_Favorite_Users_old-1)
		db.Table("users").Where("Id = ?", userId).Update("Favorite_Count", Favorite_Count_Users_old-1)
	} else {
		return false
	}
	return true
}
