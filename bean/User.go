package bean

// User 用户对象
type User struct {
	AllowAllActMsg   bool   `json:"allow_all_act_msg"`
	AllowAllComment  bool   `json:"allow_all_comment"`
	AvatarHd         string `json:"avatar_hd"`
	AvatarLarge      string `json:"avatar_large"`
	BadgeTop         string `json:"badge_top"`
	BiFollowersCount int    `json:"bi_followers_count"`
	BlockApp         int    `json:"block_app"`
	BlockWord        int    `json:"block_word"`
	City             string `json:"city"`
	Class            int    `json:"class"`
	CreatedAt        string `json:"created_at"`
	CreditScore      int    `json:"credit_score"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Extend           struct {
		Mbprivilege string `json:"mbprivilege"`
		Privacy     struct {
			Mobile int `json:"mobile"`
		} `json:"privacy"`
	} `json:"extend"`
	FavouritesCount   int    `json:"favourites_count"`
	FollowMe          bool   `json:"follow_me"`
	FollowersCount    int    `json:"followers_count"`
	Following         bool   `json:"following"`
	FriendsCount      int    `json:"friends_count"`
	Gender            string `json:"gender"`
	GeoEnabled        bool   `json:"geo_enabled"`
	HasAbilityTag     int    `json:"has_ability_tag"`
	ID                int    `json:"id"`
	Idstr             string `json:"idstr"`
	Lang              string `json:"lang"`
	Level             int    `json:"level"`
	Location          string `json:"location"`
	Mbrank            int    `json:"mbrank"`
	Mbtype            int    `json:"mbtype"`
	Name              string `json:"name"`
	OnlineStatus      int    `json:"online_status"`
	PagefriendsCount  int    `json:"pagefriends_count"`
	ProfileImageURL   string `json:"profile_image_url"`
	ProfileURL        string `json:"profile_url"`
	Province          string `json:"province"`
	Ptype             int    `json:"ptype"`
	Remark            string `json:"remark"`
	ScreenName        string `json:"screen_name"`
	Star              int    `json:"star"`
	StatusesCount     int    `json:"statuses_count"`
	Type              int    `json:"type"`
	Ulevel            int    `json:"ulevel"`
	Urank             int    `json:"urank"`
	URL               string `json:"url"`
	UserAbility       int    `json:"user_ability"`
	Verified          bool   `json:"verified"`
	VerifiedReason    string `json:"verified_reason"`
	VerifiedReasonURL string `json:"verified_reason_url"`
	VerifiedSource    string `json:"verified_source"`
	VerifiedSourceURL string `json:"verified_source_url"`
	VerifiedTrade     string `json:"verified_trade"`
	VerifiedType      int    `json:"verified_type"`
	Weihao            string `json:"weihao"`
	Status            Status `json:"status"`
}
