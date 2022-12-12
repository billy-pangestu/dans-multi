package viewmodel

// UserVM ...
type UserVM struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`

	Role RoleVM `json:"role"`
}
