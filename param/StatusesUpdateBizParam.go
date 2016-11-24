package param

// StatusesUpdateBizParam : 发布一条微博信息 参数
type StatusesUpdateBizParam struct {
	Status      string
	Visible     int32
	List_id     string
	Lat         float32
	Long        float32
	Annotations string
	Rip         string
	Is_longtext int32
}
