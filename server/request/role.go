package request

//RoleRequest ...
type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}
