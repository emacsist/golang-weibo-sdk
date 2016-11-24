package resp

import (
	"github.com/emacsist/golang-weibo-sdk/bean"
)

// StatusAndTotalNumberResp : 获取@某人的微博的结果的对象
type StatusAndTotalNumberResp struct {
	Statuses      []bean.Status `json:"statuses"`
	TotalNumber   int64         `json:"total_number"`
	PeviousCursor int64         `json:"previous_cursor"`
	NextCursor    int64         `json:"next_cursor"`
}
