package handler

import (
	"be-user-scheme/usecase"
	"net/http"
)

// RoleHandler ...
type RoleHandler struct {
	Handler
}

// FindAllHandler ...
func (h *RoleHandler) FindAllHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	roleUc := usecase.RoleUC{ContractUC: h.ContractUC}
	res, err := roleUc.FindAll()
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}
