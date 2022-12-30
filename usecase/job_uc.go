package usecase

import (
	"be-user-scheme/pkg/external"
	"be-user-scheme/pkg/logruslogger"
)

// JobUC ...
type JobUC struct {
	*ContractUC
}

// FindAll ...
func (uc JobUC) FindAll(page int, location, description, fullTime string) (res interface{}, err error) {
	ctx := "JobUC.FindAll"

	res, err = external.GetAll(page, location, description, fullTime)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_all", uc.ReqID)
		return res, err
	}
	return res, err
}

// FindByID ...
func (uc JobUC) FindByID(id string) (res interface{}, err error) {
	ctx := "JobUC.FindByID"

	res, err = external.GetByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_by_id", uc.ReqID)
		return res, err
	}
	return res, err
}
