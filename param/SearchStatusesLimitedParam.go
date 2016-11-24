package param

// SearchStatusesLimitedParam : 关键词搜索的参数。 http://open.weibo.com/wiki/C/2/search/statuses/limited
type SearchStatusesLimitedParam struct {
	Q         string
	Ids       string
	Province  string
	City      int
	Sort      int
	Starttime int64
	Endtime   int64
	Hasori    int
	Hasret    int
	Hastext   int
	Haspic    int
	Hasvideo  int
	Hasmusic  int
	Haslink   int
	Hasat     int
	Hasv      int
	Hstag     int
	Hnlynum   int
	Dup       int
	Antispam  int
	Page      int
	Count     int
	Base_app  int
}
