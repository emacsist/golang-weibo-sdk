package param

// StatusesUserTimelineBatchParam ： 批量获取用户个人微博列表参数对象
type StatusesUserTimelineBatchParam struct {
	Uids     string
	Count    int32
	Page     int32
	Flag     int32
	Base_app int32
	Feature  int32
}
