package Common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Title         string `json:"title"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id          int64  `json:"id,omitempty"`
	User        User   `json:"user"`
	Content     string `json:"content,omitempty"`
	Create_Date string `json:"create_date,omitempty"`
}

type User struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Total_Favorite int64  `json:"total_favorited,omitempty"`
	Favorite_Count int64  `json:"favorite_count,omitempty"`
	Follow_Count   int64  `json:"follow_count,omitempty"`
	Follower_Count int64  `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow,omitempty"`
}
