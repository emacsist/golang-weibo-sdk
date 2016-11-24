package param

// CommentsToMeOtherParam ： 获取某个用户收到的评论列表
type CommentsToMeOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Filter_by_source int32
	Filter_by_author int32
}
