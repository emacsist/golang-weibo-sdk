package param

// CommentsMentionsBizParam : 获取@当前登录用户的评论 api 参数
type CommentsMentionsBizParam struct {
	SinceMaxPageCountParam
	Filter_by_author int32
	Filter_by_source int32
}
