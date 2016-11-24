package param

// StatusesRepostBizParam : 转发一条微博信息 参数
type StatusesRepostBizParam struct {
	Id         int64
	Status     string
	Is_comment int32
	Rip        string
}
