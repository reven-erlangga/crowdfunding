package requests

type RegisterUserRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required" form:"email"`
	Password   string `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required" form:"email"`
	Password string `json:"password" binding:"required"`
}
