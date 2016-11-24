package param

// CommentsMentionsOtherParam ： 获取@某人的评论参数
type CommentsMentionsOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Filter_by_source int32
	Filter_by_author int32
}
