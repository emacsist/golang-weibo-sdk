package param

// CommentsShowBatchBizParam : 批量获取评论内容
type CommentsShowBatchBizParam struct {
	//需要查询的批量评论ID，用半角逗号分隔，最大50
	Cids string
}
