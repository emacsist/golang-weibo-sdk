package resp

import "github.com/emacsist/golang-weibo-sdk/bean"

// CommentsShowBatchBizResp : 批量获取评论内容，返回的是数组
type CommentsShowBatchBizResp struct {
	bean.Comment
}
