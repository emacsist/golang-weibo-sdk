package param

// StatusesShowBatchBizParam : 根据微博ID批量获取微博信息 参数
type StatusesShowBatchBizParam struct {
	Ids           string
	Trim_user     int32
	IsGetLongText int32
}
