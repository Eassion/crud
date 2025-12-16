package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"gte=0,lte=150"`
}

type UpdateUserRequest struct {
	Age int `json:"age" binding:"required,gte=0,lte=150"`
}
