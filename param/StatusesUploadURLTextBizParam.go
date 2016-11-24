package param

// StatusesUploadURLTextBizParam ： 参数
// picid或url必选一个，两个参数都存在时，取picid参数的值为准；
type StatusesUploadURLTextBizParam struct {
	Status      string
	Visible     int32
	List_id     string
	Url         string
	Pic_id      string
	Lat         string
	Long        float32
	Annotations string
	Rip         string
}
