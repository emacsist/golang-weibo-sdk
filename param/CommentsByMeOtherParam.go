package param

// CommentsByMeOtherParam ： 获取某个用户发出的评论列表
type CommentsByMeOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Filter_by_source int32
}
