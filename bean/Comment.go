package bean

// Comment : 微博评论对象
type Comment struct {
	CreatedAt    string   `json:"created_at"`
	ID           int64    `json:"id"`
	Text         string   `json:"text"`
	Source       string   `json:"source"`
	Mid          string   `json:"mid"`
	User         User     `json:"user"`
	Status       Status   `json:"status"`
	ReplyComment *Comment `json:"reply_comment"`
}
