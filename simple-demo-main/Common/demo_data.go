package Common

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:          1,
		User:        DemoUser,
		Content:     "Test Comment",
		Create_Date: "05-01",
	},
}

var DemoUser = User{
	Id:             1,
	Name:           "TestUser",
	Follow_Count:   0,
	Follower_Count: 0,
	IsFollow:       false,
}
