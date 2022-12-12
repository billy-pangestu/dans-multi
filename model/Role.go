package model

import (
	"be-user-scheme/helper"
	"be-user-scheme/pkg/pg"
	"database/sql"
)

var (
	// RoleCodeSuperadmin ...
	RoleCodeSuperadmin = "superadmin"
	// RoleCodeAdmin ...
	RoleCodeAdmin = "admin"
)

// IRole ...
type IRole interface {
	FindAll() ([]RoleEntity, error)
	FindByID(id string) (RoleEntity, error)
	FindByCode(code string) (RoleEntity, error)
}

// RoleEntity ....
type RoleEntity struct {
	ID        string         `gorm:"id"`
	Name      sql.NullString `gorm:"name"`
	CreatedAt string         `gorm:"created_at"`
	UpdatedAt string         `gorm:"updated_at"`
	DeletedAt sql.NullString `gorm:"deleted_at"`
}

func (RoleEntity) TableName() string {
	return helper.UserRoleDBName
}

// roleModel ...
type roleModel struct {
	DB *pg.MySQL
}

// NewRoleModel ...
func NewRoleModel(db *pg.MySQL) IRole {
	return &roleModel{DB: db}
}

// FindAll ...
func (model roleModel) FindAll() (data []RoleEntity, err error) {
	err = model.DB.Find(&data).Error

	return data, err
}

// FindByID ...
func (model roleModel) FindByID(id string) (data RoleEntity, err error) {
	query := `"id" = $1`
	err = model.DB.Where(query, id).Find(&data).Error

	return data, err
}

// FindByCode ...
func (model roleModel) FindByCode(code string) (data RoleEntity, err error) {
	// query :=
	// 	`SELECT "id", "code", "name", "status", "created_at", "updated_at", "deleted_at"
	// 	FROM "roles" WHERE "deleted_at" IS NULL AND "code" = $1
	// 	ORDER BY "created_at" DESC LIMIT 1`
	// err = model.DB.QueryRow(query, code).Scan(
	// 	&data.ID, &data.Code, &data.Name, &data.Status, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
	// )

	return data, err
}
