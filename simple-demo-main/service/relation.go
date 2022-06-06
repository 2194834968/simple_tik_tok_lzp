package service

import "github.com/RaymondCode/simple-demo/Common"

type Follower struct {
	Id             int64
	Be_Follower_Id int64
	Follower_Id    int64
}

//userId对toUserId执行关注/取关
func RelationAction(userId int64, toUserId int64, actionType int64) bool {
	db := Common.MysqlDb

	var newFollower Follower
	var toUserId_followerCount int64
	var userId_followCount int64

	db.Table("users").Select("follower_Count").Where("Id = ?", toUserId).Take(&toUserId_followerCount)
	db.Table("users").Select("follow_Count").Where("Id = ?", userId).Take(&userId_followCount)

	if actionType == 1 {
		newFollower = Follower{
			Be_Follower_Id: toUserId,
			Follower_Id:    userId,
		}
		db.Table("follower").Create(&newFollower)
		db.Table("users").Where("Id = ?", toUserId).Update("follower_Count", toUserId_followerCount+1)
		db.Table("users").Where("Id = ?", userId).Update("follow_Count", userId_followCount+1)
	} else if actionType == 2 {
		db.Table("follower").Delete(&newFollower)
		db.Table("users").Where("Id = ?", toUserId).Update("follower_Count", toUserId_followerCount-1)
		db.Table("users").Where("Id = ?", userId).Update("follow_Count", userId_followCount-1)
	} else {
		return false
	}

	return true
}

//获取userId关注的人
func RelationFollowList(userId int64, myUserId int64) []Common.User {
	db := Common.MysqlDb

	var FollowIdList []int64

	//用于返回的关注列表
	var FollowList []Common.User

	db.Table("follower").Select("be_follower_id").Where("Follower_Id = ?", userId).Find(&FollowIdList)

	for i := 0; i < len(FollowIdList); i++ {
		var BeFollower Common.User
		BeFollower = UserInfo(FollowIdList[i], myUserId)
		FollowList = append(FollowList, BeFollower)
	}

	return FollowList
}

//获取关注userId的人
func RelationFollowerList(userId int64, myUserId int64) []Common.User {
	db := Common.MysqlDb

	var FollowerIdList []int64

	//用于返回的关注列表
	var FollowerList []Common.User

	db.Table("follower").Select("follower_id").Where("be_follower_Id = ?", userId).Find(&FollowerIdList)

	for i := 0; i < len(FollowerIdList); i++ {
		var BeFollower Common.User
		BeFollower = UserInfo(FollowerIdList[i], myUserId)
		FollowerList = append(FollowerList, BeFollower)
	}

	return FollowerList
}
