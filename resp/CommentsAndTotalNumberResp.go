package resp

import (
	"github.com/emacsist/golang-weibo-sdk/bean"
)

// CommentsAndTotalNumberResp : 评论与总数的结果对象
type CommentsAndTotalNumberResp struct {
	Comments      []bean.Comment `json:"comments"`
	TotalNumber   int64          `json:"total_number"`
	PeviousCursor int64          `json:"previous_cursor"`
	NextCursor    int64          `json:"next_cursor"`
}
