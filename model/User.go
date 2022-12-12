package model

import (
	"database/sql"
	"skeleton-backend/helper"
	"skeleton-backend/pkg/pg"
)

// IUser ...
type IUser interface {
	FindAll() ([]UserEntity, error)
	FindByUsername(username string) (UserEntity, error)
	FindByID(id string) (UserEntity, error)
	Store(data UserEntity) (UserEntity, error)
	Update(data UserEntity) error
	Delete(data UserEntity) error
}

// UserEntity ....
type UserEntity struct {
	ID        string         `gorm:"id"`
	FirstName sql.NullString `gorm:"first_name"`
	LastName  sql.NullString `gorm:"last_name"`
	Username  sql.NullString `gorm:"username"`
	Password  sql.NullString `gorm:"password"`
	CreatedAt sql.NullString `gorm:"createdAt"`
	UpdatedAt sql.NullString `gorm:"updatedAt"`
	DeletedAt sql.NullString `gorm:"deletedAt"`
	RoleID    string         `gorm:"role_id"`
}

func (UserEntity) TableName() string {
	return helper.UserDBName
}

// userModel ...
type userModel struct {
	DB *pg.MySQL
}

// NewUserModel ...
func NewUserModel(mysql *pg.MySQL) IUser {
	return &userModel{
		DB: mysql,
	}
}

// FindByID ...
func (model userModel) FindAll() (data []UserEntity, err error) {
	err = model.DB.Find(&data).Error
	return
}

// FindByID ...
func (model userModel) FindByID(id string) (data UserEntity, err error) {
	query := `"id"=$1`

	err = model.DB.Where(query, id).Find(&data).Error

	return
}

// FindByUsername ...
func (model userModel) FindByUsername(username string) (data UserEntity, err error) {
	query := `"username"=$1`

	err = model.DB.Where(query, username).Find(&data).Joins("Roles").Error

	return
}

// Store ...
func (model userModel) Store(data UserEntity) (res UserEntity, err error) {
	err = model.DB.Save(&data).Find(&res).Error

	return
}

// Update ...
func (model userModel) Update(data UserEntity) (err error) {
	err = model.DB.Model(&data).Update(&data).Error

	return
}

// Delete ...
func (model userModel) Delete(data UserEntity) (err error){
	err = model.DB.Model(&data).Update(&data).Error
	
	return
}
