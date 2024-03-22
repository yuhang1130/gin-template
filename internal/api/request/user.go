package request

type UserLoginDto struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"passWord"  binding:"required"`
}

type UserCreateDto struct {
	UserLoginDto
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email"  binding:"required"`
}

type UserPageQueryDTO struct {
	Name     string `json:"name"`     // 分页查询的name
	Page     int    `json:"page"`     // 分页查询的页数
	PageSize int    `json:"pageSize"` // 分页查询的页容量
}
