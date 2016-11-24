package param

// CommentsTimelineOtherParam ： 获取某个用户发出和收到的评论列表参数
type CommentsTimelineOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Trim_user int32
}
