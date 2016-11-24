package bean

// Attitude : 收到的赞的对象
type Attitude struct {
	ID           int64  `json:"id"`
	CreatedAt    string `json:"created_at"`
	Attitude     string `json:"attitude"`
	LastAttitude string `json:"last_attitude"`
	Source       string `json:"source"`
	User         User   `json:"user"`
	Status       Status `json:"status"`
}
