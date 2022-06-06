package service

import (
	"github.com/RaymondCode/simple-demo/Common"
	"time"
)

type CommentDataBase struct {
	Id          int64
	Video_Id    int64
	User_Id     int64
	Content     string
	Create_Date time.Time
}

func CommentAction(videoId int64, userId int64, action_type int64, comment_text string, comment_id int64) Common.Comment {
	db := Common.MysqlDb

	var CommentCount int64
	db.Table("videos").Select("Comment_Count").Where("Id = ?", videoId).Take(&CommentCount)

	//用于数据库查询的评论
	var commentDataBase CommentDataBase

	//用于返回的评论
	var comment Common.Comment

	if action_type == 1 {
		//用于数据库查询的评论
		commentDataBase = CommentDataBase{
			Video_Id:    videoId,
			User_Id:     userId,
			Content:     comment_text,
			Create_Date: time.Now(),
		}

		db.Table("comment").Create(&commentDataBase)
		db.Table("videos").Where("Id = ?", videoId).Update("Comment_Count", CommentCount+1)

		var CommentIdNew int64
		db.Table("comment").Select("Id").Where("video_id = ? AND user_id = ?", videoId, userId).Take(&CommentIdNew)

		//用于返回的评论
		comment = Common.Comment{
			Id:          CommentIdNew,
			User:        UserInfo(userId, userId),
			Content:     comment_text,
			Create_Date: commentDataBase.Create_Date.Format("2006-01-02"),
		}

	} else if action_type == 2 {
		db.Table("comment").Where("Id = ?", comment_id).Delete(&commentDataBase)
		db.Table("videos").Where("Id = ?", videoId).Update("Comment_Count", CommentCount-1)
	}
	return comment
}

func CommentList(videoId int64, myUserId int64) []Common.Comment {
	db := Common.MysqlDb

	//用于返回的评论列表
	var commentList []Common.Comment

	//用于数据库查询的评论列表
	var commentListDataBase []CommentDataBase

	db.Table("comment").Where("video_id = ?", videoId).Find(&commentListDataBase)

	for i := 0; i < len(commentListDataBase); i++ {
		var commentTemp Common.Comment

		commentTemp.Id = commentListDataBase[i].Id
		commentTemp.User = UserInfo(commentListDataBase[i].User_Id, myUserId)
		commentTemp.Content = commentListDataBase[i].Content
		commentTemp.Create_Date = commentListDataBase[i].Create_Date.Format("2006-01-02")

		commentList = append(commentList, commentTemp)
	}

	return commentList
}
