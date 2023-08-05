package model

type Users struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type UpdateUserRequest struct {
	Id           int `validate:required"`
	UserStandard UserRequest
}

type UserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserStandard UserRequest
}

type LoginResponse struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"token"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
