package param

import (
	"encoding/json"
	"os"
)

// StatusesUploadBizParam : 上传图片并发布一条微博 参数
// 还有一个pic字段，这个额外传，因为它是一个文件
type StatusesUploadBizParam struct {
	Status      string
	Visible     int32
	List_id     string
	Lat         string
	Long        float32
	Annotations string
	Rip         string
	Pic         *os.File
}

func (s *StatusesUploadBizParam) String() string {
	body, _ := json.Marshal(s)
	return string(body)
}
