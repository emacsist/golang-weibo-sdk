package resp

import (
	"github.com/emacsist/golang-weibo-sdk/bean"
)

// UsersAndTotalNumberResp : 用户和总数的结果对象
type UsersAndTotalNumberResp struct {
	Users         []bean.User `json:"users"`
	TotalNumber   int64       `json:"total_number"`
	PeviousCursor int64       `json:"previous_cursor"`
	NextCursor    int64       `json:"next_cursor"`
}
