package param

// CommentsByMeBizParam ： 获取当前登录用户发出的评论列表 参数
type CommentsByMeBizParam struct {
	SinceMaxPageCountParam
	Filter_by_source int32
}
