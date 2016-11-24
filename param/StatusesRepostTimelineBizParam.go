package param

// StatusesRepostTimelineBizParam ： 获取当前登录用户发布的一条原创微博的最新转发微博
type StatusesRepostTimelineBizParam struct {
	Id int64
	SinceMaxPageCountParam
	Filter_by_author int8
}
