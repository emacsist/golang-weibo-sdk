package param

// CommentsReplyBizParam : 回复一条评论
type CommentsReplyBizParam struct {
	Cid             int64
	Id              int64
	Comment         string
	Without_mention int32
	Comment_ori     int32
	Rip             string
}
