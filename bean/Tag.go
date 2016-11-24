package bean

// Tag ： 标签对象
type Tag struct {
	ID int64 `json:"id"`

	// Tags，它返回的是一个map，然后自己去遍历它，因为它的Key是动态的
	Tags interface{} `json:"tags"`
}
