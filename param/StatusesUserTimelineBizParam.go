package param

// StatusesUserTimelineBizParam ： 获取当前登录用户发布的微博 参数对象
type StatusesUserTimelineBizParam struct {
	SinceMaxPageCountParam
	Flag      int32
	Base_app  int32
	Feature   int32
	Trim_user int32
}
