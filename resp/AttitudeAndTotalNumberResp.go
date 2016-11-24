package resp

import "github.com/emacsist/golang-weibo-sdk/bean"

// AttitudeAndTotalNumberResp : 赞列表和总数
type AttitudeAndTotalNumberResp struct {
	Attitudes     []bean.Attitude `json:"attitudes"`
	TotalNumber   int64           `json:"total_number"`
	PeviousCursor int64           `json:"previous_cursor"`
	NextCursor    int64           `json:"next_cursor"`
}
