package usecase

import (
	"database/sql"
	"errors"
	"skeleton-backend/helper"
	"skeleton-backend/model"
	"skeleton-backend/pkg/logruslogger"
	"skeleton-backend/server/request"
	"skeleton-backend/usecase/viewmodel"
	"time"
)

// UserUC ...
type UserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserUC) BuildBody(data *model.UserEntity, res *viewmodel.UserVM, showPass bool, userRole *viewmodel.RoleVM) {
	res.ID = data.ID
	res.FirstName = data.FirstName.String
	res.LastName = data.LastName.String
	res.Username = data.Username.String
	res.CreatedAt = data.CreatedAt.String
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	if showPass {
		res.Password = data.Password.String
	}

	res.Role = *userRole
}

// FindAll ...
func (uc UserUC) FindAll() (res []viewmodel.UserVM, err error) {
	ctx := "UserUC.FindAll"

	userModel := model.NewUserModel(uc.DB)
	data, err := userModel.FindAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	roleUc := RoleUC{ContractUC: uc.ContractUC}
	userRoles, err := roleUc.FindAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_role_not_found", uc.ReqID)
		return res, err
	}

	for _, datum := range data {
		var tempUser viewmodel.UserVM
		var tempRole viewmodel.RoleVM

		for _, userRole := range userRoles {
			if datum.RoleID == userRole.ID {
				tempRole = userRole
			}
		}

		uc.BuildBody(&datum, &tempUser, false, &tempRole)

		res = append(res, tempUser)
	}

	return res, err
}

// FindByID ...
func (uc UserUC) FindByID(id string, showPass bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByID"

	userModel := model.NewUserModel(uc.DB)
	data, err := userModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	roleUc := RoleUC{ContractUC: uc.ContractUC}
	userRole, err := roleUc.FindByID(data.RoleID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_role_not_found", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showPass, &userRole)

	return res, err
}

// FindByUsername ...
func (uc UserUC) FindByUsername(username string, showPass bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByUsername"
	userModel := model.NewUserModel(uc.DB)

	data, err := userModel.FindByUsername(username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, err
	}

	roleUc := RoleUC{ContractUC: uc.ContractUC}
	userRole, err := roleUc.FindByID(data.RoleID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_role_not_found", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showPass, &userRole)

	return res, err
}

// Create ...
func (uc UserUC) Create(data request.UserRequest) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.Create"

	// Decrypt password dari frontend karena frontend implementasi aes untuk menencrypsi password yang akan dikirimkan
	// authUc := AuthUC{ContractUC: uc.ContractUC}
	// passwordInput, err := authUc.DecryptedOnly(data.Password)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	//password diubah dalam bentuk hash
	password, err := helper.HashPassword(data.Password)
	if err != nil {
		return res, err
	}

	//Error Checking Double Username
	userChecking, _ := uc.FindByUsername(data.Username, false)
	if userChecking.ID != "" {
		err = errors.New("duplicate username")
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_user_by_username", uc.ReqID)
		return res, err
	}

	// Error Checking Role ID
	roleUc := RoleUC{ContractUC: uc.ContractUC}
	resRole, err := roleUc.FindByID(data.RoleID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_role_by_id", uc.ReqID)
		return res, errors.New("role not found")
	}

	now := time.Now().UTC()

	dbInput := model.UserEntity{
		FirstName: sql.NullString{data.FirstName, true},
		LastName:  sql.NullString{data.LastName, true},
		Username:  sql.NullString{data.Username, true},
		Password:  sql.NullString{password, true},
		CreatedAt: sql.NullString{now.Format(time.RFC3339Nano), true},
		UpdatedAt: sql.NullString{now.Format(time.RFC3339Nano), true},
		RoleID:    data.RoleID,
	}

	userModel := model.NewUserModel(uc.DB)
	dbRes, err := userModel.Store(dbInput)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&dbRes, &res, false, &resRole)

	return res, err
}

// Update ...
func (uc UserUC) Update(data request.UserUpdateRequest) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.Update"

	// Error Checking Role ID
	roleUc := RoleUC{ContractUC: uc.ContractUC}
	_, err = roleUc.FindByID(data.RoleID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_role_by_id", uc.ReqID)
		return res, errors.New("role not found")
	}

	now := time.Now().UTC()

	dbInput := model.UserEntity{
		ID:        data.ID,
		FirstName: sql.NullString{data.FirstName, true},
		LastName:  sql.NullString{data.LastName, true},
		UpdatedAt: sql.NullString{now.Format(time.RFC3339Nano), true},
		RoleID:    data.RoleID,
	}

	userModel := model.NewUserModel(uc.DB)
	err = userModel.Update(dbInput)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	res, err = uc.FindByID(dbInput.ID, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_find_id", uc.ReqID)
		return res, err
	}

	return res, err
}

// Delete ...
func (uc UserUC) Delete(data request.UserDeleteRequest) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.Delete"

	now := time.Now().UTC()

	dbInput := model.UserEntity{
		ID:        data.ID,
		DeletedAt: sql.NullString{now.Format(time.RFC3339Nano), true},
	}

	userModel := model.NewUserModel(uc.DB)
	err = userModel.Update(dbInput)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	return res, err
}

// Login ...
func (uc UserUC) Login(data request.UserLoginRequest) (res viewmodel.JwtVM, err error) {
	ctx := "UserUC.Login"

	user, err := uc.FindByUsername(data.Username, true)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.Username, ctx, "user_not_found", uc.ReqID)
		return res, errors.New("user not found")
	}

	// Uncomment this if Frontend Implement AES on Password
	// authUc := AuthUC{ContractUC: uc.ContractUC}
	// passwordInput, err := authUc.DecryptedOnly(data.Password)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	passwordInput := data.Password

	// password checking using hash
	// note: password saved on db using hash
	match := helper.CheckPasswordHash(passwordInput, user.Password)
	if !match {
		logruslogger.Log(logruslogger.WarnLevel, "invalid_password", ctx, "invalid_password", uc.ReqID)
		return res, errors.New(helper.InvalidCredentials)
	}

	// Jwe the payload & Generate jwt token
	payload := map[string]interface{}{
		"id":        user.ID,
		"role_name": user.Role.Name,
		"unique_id": user.Username,
	}

	// generate jwt token and store to redis
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwe", uc.ReqID)
		return res, errors.New(helper.InternalServer)
	}

	return res, err
}

// Logout ...
func (uc UserUC) Logout(token, userID string) (res viewmodel.JwtVM, err error) {

	err = uc.RemoveFromRedis("userDeviceID" + userID)
	if err != nil {
		return res, err
	}

	return res, err
}
