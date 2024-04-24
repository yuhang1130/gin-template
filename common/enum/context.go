package enum

const (
	CurrentId   = "currentId"
	CurrentName = "currentName"
)

const (
	MaxPageSize int = 100 // 单页最大数量
	MinPageSize int = 10  // 单页最小数量

	Sid         string = "gin.sid"
	SessionID   string = "sessionID"
	SessionData string = "sessionData"
)

type PartialUser struct {
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	IsAdmin  bool   `json:"isAdmin"`
	Email    string `json:"email"`
}

type SessionDto struct {
	UserId   uint // 默认项目ID，如果要获取当前项目ID, 请使用OpUserId ProjectId
	OpUserId uint //  操作者ID，只有request请求时会有这ID。表示请求是以这个UserId为准，例如CreateUserId的过滤，或者创建的时候身份人
	Rights   []string
	User     PartialUser
}
