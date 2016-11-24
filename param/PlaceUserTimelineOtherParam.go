package param

// PlaceUserTimelineOtherParam : 获取某个用户的位置动态
type PlaceUserTimelineOtherParam struct {
	Uid int64
	SinceMaxPageCountParam
	Base_app int32
}
