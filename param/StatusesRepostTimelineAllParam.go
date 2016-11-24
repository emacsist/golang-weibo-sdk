package param

// StatusesRepostTimelineAllParam ： 返回一条微博的全部转发微博列表参数对象
type StatusesRepostTimelineAllParam struct {
	Id               int64
	Count            int32
	Page             int32
	Filter_by_author int8
}