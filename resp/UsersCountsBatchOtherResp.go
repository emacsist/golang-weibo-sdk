package resp

// UsersCountsBatchOtherResp : 批量获取用户的粉丝数、关注数、微博数
type UsersCountsBatchOtherResp struct {
	ID                string `json:"id"`
	FollowersCount    string `json:"followers_count"`
	FriendsCount      string `json:"friends_count"`
	MutualFriendCount string `json:"mutual_friend_count"`
	StatusesCount     string `json:"statuses_count"`
}
