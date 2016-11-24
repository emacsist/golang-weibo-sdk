package param

// CommentsShowAllParam ： 批量获取用户个人微博列表参数对象
type CommentsShowAllParam struct {
	Id               string
	Count            int32
	Page             int32
	Filter_by_author int32
}
