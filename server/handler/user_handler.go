package handler

import (
	"net/http"
	"strings"

	"skeleton-backend/model"
	"skeleton-backend/pkg/str"
	"skeleton-backend/server/request"
	"skeleton-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserHandler ...
type UserHandler struct {
	Handler
}

// LoginHandler ...
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	req := request.UserLoginRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {
		SendBadRequest(w, "Transaction")
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Login(req)
	if err != nil {

		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}

	h.ContractUC.Tx.Commit()
	SendSuccess(w, res, nil)
	return
}

// LogoutHandler ...
func (h *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	tokenAuthHeader := r.Header.Get("Authorization")
	tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {
		SendBadRequest(w, "Transaction")
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Logout(tokenAuth, userID)
	if err != nil {
		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}
	h.ContractUC.Tx.Commit()

	SendSuccess(w, res, nil)
	return
}

// TokenHandler ...
func (h *UserHandler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {
		SendBadRequest(w, "Transaction")
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.FindByID(userID, false)
	if err != nil {
		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}
	h.ContractUC.Tx.Commit()

	SendSuccess(w, res, nil)
	return
}

// CreateHandler ...
func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	req := request.UserRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {

		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {

		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	req.UniqueID = str.RandomNumericString(10)

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {

		SendBadRequest(w, "Transaction")
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Create(req)
	if err != nil {

		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}

	h.ContractUC.Tx.Commit()
	SendSuccess(w, res, nil)
	return
}

// AddFund ...
func (h *UserHandler) AddFundHandler(w http.ResponseWriter, r *http.Request) {

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	req := request.UserAddFundRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {

		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {

		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {

		SendBadRequest(w, "Transaction")
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.AddFund(userID, req)
	if err != nil {

		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}

	h.ContractUC.Tx.Commit()
	SendSuccess(w, res, nil)
	return
}