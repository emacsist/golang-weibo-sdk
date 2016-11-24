package param

// StatusesMentionsOtherParam :  获取@某人的微博的参数对象
type StatusesMentionsOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Filter_by_author int8
	Filter_by_source int8
	Filter_by_type   int8
}
