package resp

import (
	"github.com/emacsist/golang-weibo-sdk/bean"
)

// TagsTagsBatchOtherResp :批量获取用户标签结果对象
/*
[{"id":5225532117,"tags":[{"201103260002460337":"美图摄影","weight":"7069191","flag":"0"}]}]
*/
type TagsTagsBatchOtherResp struct {
	bean.Tag
}
