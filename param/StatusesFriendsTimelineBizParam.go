package param

// StatusesFriendsTimelineBizParam :  获取当前登录用户及其所关注用户的最新微博 参数
type StatusesFriendsTimelineBizParam struct {
	SinceMaxPageCountParam
	Base_app  int32
	Feature   int32
	Trim_user int32
}
