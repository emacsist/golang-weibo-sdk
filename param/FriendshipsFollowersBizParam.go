package param

// FriendshipsFollowersBizParam ： 获取当前登录用户的粉丝列表
type FriendshipsFollowersBizParam struct {
	Gender   string
	Province int32
	City     int32
	// 指定接收用户的年龄范围，取值内容为：6-12、13-18、19-22、23-25、26-29、30-39、40-59、60-80、其他。
	Age    string
	Type   int32
	Flag   int32
	Bucket int32
	// 返回结果的时间戳游标，若指定此参数，则返回关注时间小于或等于max_time的粉丝，默认从当前时间开始算。返回结果中会得到next_cursor字段，表示下一页的max_time。next_cursor为0表示已经到记录末尾
	Max_Time   string
	Cursor_uid int64
}
