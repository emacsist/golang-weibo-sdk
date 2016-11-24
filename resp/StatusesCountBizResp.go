package resp

// StatusesCountBizResp : 批量获取指定微博的转发数评论数喜欢数 响应对象
// 返回的结果是这个类型的数组
/*
[
    {
        "id": "32817222",
        "comments": "16",
        "reposts": "38",
        "likes": "38"
    },
   ...
]
*/
type StatusesCountBizResp struct {
	ID       string `json:"id"`
	Comments string `json:"comments"`
	Reposts  string `json:"reposts"`
	Likes    string `json:"likes"`
}
