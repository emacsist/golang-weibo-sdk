package resp

import (
	"encoding/json"
)

// TagsBizResp : 当前登录用户 的标签
type TagsBizResp struct {
	Body string
}

// GetTags 将返回的结果转换为map数组
func (t TagsBizResp) GetTags() (resp []map[string]interface{}) {
	json.Unmarshal([]byte(t.Body), &resp)
	return
}
