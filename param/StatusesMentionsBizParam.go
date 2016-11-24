package param

// StatusesMentionsBizParam :  获取@当前登录用户的微博列表 参数
type StatusesMentionsBizParam struct {
	SinceMaxPageCountParam
	Filter_by_author int8
	Filter_by_source int8
	Filter_by_type   int8
}
