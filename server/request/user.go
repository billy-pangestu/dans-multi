package request

//UserRequest ...
type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	RoleID    string `json:"role_id" validate:"required"`
}

//UserUpdateRequest ...
type UserUpdateRequest struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	RoleID    string `json:"role_id" validate:"required"`
}

//UserDeleteRequest ...
type UserDeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

//UserLoginRequest ...
type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}

//UserAddFundRequest ...
type UserAddFundRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
