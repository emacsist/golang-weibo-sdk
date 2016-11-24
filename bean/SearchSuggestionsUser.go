package bean

// SearchSuggestionsUser : 搜用户搜索建议
type SearchSuggestionsUser struct {
	ScreenName     string `json:"screen_name"`
	FollowersCount int    `json:"followers_count"`
	UID            int    `json:"uid"`
}
