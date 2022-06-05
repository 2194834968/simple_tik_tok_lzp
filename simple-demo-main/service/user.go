package service

import (
	"github.com/RaymondCode/simple-demo/Common"
	"strings"
	"time"
)

//用于数据库查询的用户结构体
type UserDatabase struct {
	Id             int64
	Name           string
	Email          string
	Password       string
	Total_Favorite int64
	Favorite_Count int64
	Follow_Count   int64
	Follower_Count int64
	Created_At     time.Time
	Updated_At     time.Time
}

func Register(Email string, Password string) (int64, bool) {
	db := Common.MysqlConnection()

	//判断userName是否存在
	var EmailExist string
	db.Table("users").Select("Email").Where("Email = ?", Email).Take(&EmailExist)
	if strings.EqualFold(Email, EmailExist) {
		//fmt.Printf("存在了存在了")
		return 0, false
	}

	//用于数据库查询的用户结构体
	userDatabase := UserDatabase{
		Name:           Email,
		Email:          Email,
		Password:       Password,
		Total_Favorite: 0,
		Favorite_Count: 0,
		Follow_Count:   0,
		Follower_Count: 0,
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}

	var userid int64
	db.Table("users").Create(&userDatabase)
	db.Table("users").Select("Id").Where("Email = ?", Email).Take(&userid)
	if userid == 0 {
		return 0, false
	}
	return userid, true

}

func Login(Email string, Password string) (int64, bool) {
	db := Common.MysqlConnection()

	var userid int64

	db.Table("users").Select("Id").Where("Email = ? AND Password = ?", Email, Password).Take(&userid)
	if userid == 0 {
		//fmt.Printf("错了错了")
		return 0, false
	}
	//fmt.Printf("对了对了")
	return userid, true

}

func UserInfo(userid int64) Common.User {
	db := Common.MysqlConnection()

	var user Common.User

	db.Table("users").Where("Id = ?", userid).Take(&user)
	//fmt.Printf("\n AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFCount %d \n", user.Total_Favorite)
	user.IsFollow = false
	/*
		var userTest UserDatabase

		db.Table("users").Where("Id = ?", userid).Take(&userTest)
		fmt.Printf("\n CreateAt %s \n", userTest.Created_At.String())
		fmt.Printf("\n UpdatedAt %s \n", userTest.Updated_At.String())
	*/
	//fmt.Printf("%d", user.IsFollow)

	return user
}
