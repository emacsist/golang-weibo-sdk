package param

// CommentsShowBizParam : 获取当前登录用户发布的微博下的评论列表
type CommentsShowBizParam struct {
	Id int64
	SinceMaxPageCountParam
	Filter_by_author int32
}
