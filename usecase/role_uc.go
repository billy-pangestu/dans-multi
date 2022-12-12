package usecase

import (
	"skeleton-backend/model"
	"skeleton-backend/pkg/logruslogger"
	"skeleton-backend/usecase/viewmodel"
)

// RoleUC ...
type RoleUC struct {
	*ContractUC
}

// BuildBody ...
func (uc RoleUC) BuildBody(data *model.RoleEntity, res *viewmodel.RoleVM) {
	res.ID = data.ID
	res.Name = data.Name.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt
	res.DeletedAt = data.DeletedAt.String
}

// FindByID ...
func (uc RoleUC) FindAll() (res []viewmodel.RoleVM, err error) {
	ctx := "RoleUC.FindAll"

	roleModel := model.NewRoleModel(uc.DB)
	data, err := roleModel.FindAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	for _, datum := range data {
		var temp viewmodel.RoleVM
		uc.BuildBody(&datum, &temp)
		res = append(res, temp)
	}

	return res, err
}

// FindByID ...
func (uc RoleUC) FindByID(id string) (res viewmodel.RoleVM, err error) {
	ctx := "RoleUC.FindByID"

	roleModel := model.NewRoleModel(uc.DB)
	data, err := roleModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res)

	return res, err
}

// FindByCode ...
func (uc RoleUC) FindByCode(code string) (res viewmodel.RoleVM, err error) {
	// ctx := "RoleUC.FindByCode"

	// roleModel := model.NewRoleModel(uc.DB)
	// data, err := roleModel.FindByCode(code)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	// res = viewmodel.RoleVM{
	// 	ID:        data.ID,
	// 	Code:      data.Code.String,
	// 	Name:      data.Name.String,
	// 	Status:    data.Status.Bool,
	// 	CreatedAt: data.CreatedAt,
	// 	UpdatedAt: data.UpdatedAt,
	// 	DeletedAt: data.DeletedAt.String,
	// }

	return res, err
}
