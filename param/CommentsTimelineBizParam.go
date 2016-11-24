package param

// CommentsTimelineBizParam :   获取当前登录用户发出及收到的评论列表 参数
type CommentsTimelineBizParam struct {
	SinceMaxPageCountParam
	Trim_user int32
}
