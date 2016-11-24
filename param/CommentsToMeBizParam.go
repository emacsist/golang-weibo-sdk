package param

// CommentsToMeBizParam : 获取当前登录用户收到的评论列表参数
type CommentsToMeBizParam struct {
	SinceMaxPageCountParam
	Filter_by_author int32
	Filter_by_source int32
}
