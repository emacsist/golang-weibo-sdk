package resp

import (
	"github.com/emacsist/golang-weibo-sdk/bean"
)

// StatusesRepostAndTotalNumberResp : 转发微博的结果（数据和totalNumber)
type StatusesRepostAndTotalNumberResp struct {
	Reposts       []bean.Status `json:"reposts"`
	PeviousCursor int64         `json:"previous_cursor"`
	NextCursor    int64         `json:"next_cursor"`
	TotalNumber   int64         `json:"total_number"`
}
