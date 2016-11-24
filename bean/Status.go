package bean

// Status : 微博对象
type Status struct {
	Appid          int           `json:"appid"`
	AttitudesCount int           `json:"attitudes_count"`
	Base62ID       string        `json:"base62_id"`
	BizFeature     int           `json:"biz_feature"`
	Category       int           `json:"category"`
	CommentsCount  int           `json:"comments_count"`
	CreatedAt      string        `json:"created_at"`
	DarwinTags     []interface{} `json:"darwin_tags"`
	Favorited      bool          `json:"favorited"`
	Geo            struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geo"`
	GifIds              string        `json:"gif_ids"`
	HasActionTypeCard   int           `json:"hasActionTypeCard"`
	HotWeiboTags        []interface{} `json:"hot_weibo_tags"`
	ID                  int           `json:"id"`
	Idstr               string        `json:"idstr"`
	InReplyToScreenName string        `json:"in_reply_to_screen_name"`
	InReplyToStatusID   string        `json:"in_reply_to_status_id"`
	InReplyToUserID     string        `json:"in_reply_to_user_id"`
	IsLongText          bool          `json:"isLongText"`
	IsShowBulletin      int           `json:"is_show_bulletin"`
	Mid                 string        `json:"mid"`
	Mlevel              int           `json:"mlevel"`
	PicIds              []interface{} `json:"pic_ids"`
	PositiveRecomFlag   int           `json:"positive_recom_flag"`
	Relation            interface{}   `json:"relation"`
	RepostsCount        int           `json:"reposts_count"`
	Rid                 string        `json:"rid"`
	Source              string        `json:"source"`
	SourceAllowclick    int           `json:"source_allowclick"`
	SourceType          int           `json:"source_type"`
	Status              int           `json:"status"`
	Text                string        `json:"text"`
	TextLength          int           `json:"textLength"`
	TextTagTips         []interface{} `json:"text_tag_tips"`
	Truncated           bool          `json:"truncated"`
	URLObjects          []struct {
		AssoLikeCount       int  `json:"asso_like_count"`
		CardInfoUnIntegrity bool `json:"card_info_un_integrity"`
		FollowerCount       int  `json:"follower_count"`
		Info                struct {
			Description  string `json:"description"`
			LastModified int    `json:"last_modified"`
			Result       bool   `json:"result"`
			Title        string `json:"title"`
			Transcode    int    `json:"transcode"`
			Type         int    `json:"type"`
			URLLong      string `json:"url_long"`
			URLShort     string `json:"url_short"`
		} `json:"info"`
		IsActionType          bool   `json:"isActionType"`
		LikeCount             int    `json:"like_count"`
		SuperTopicPhotoCount  int    `json:"super_topic_photo_count"`
		SuperTopicStatusCount int    `json:"super_topic_status_count"`
		URLOri                string `json:"url_ori"`
	} `json:"url_objects"`
	User     *User `json:"user"`
	UserType int   `json:"userType"`
	Visible  struct {
		ListID int `json:"list_id"`
		Type   int `json:"type"`
	} `json:"visible"`
	RetweetedStatus *Status `json:"retweeted_status"`
}
