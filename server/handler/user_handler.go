package handler

import (
	"net/http"
	"strings"

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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Login(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// LogoutHandler ...
func (h *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	tokenAuthHeader := r.Header.Get("Authorization")
	tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Logout(tokenAuth, userID)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// TokenHandler ...
func (h *UserHandler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.FindByID(userID, false)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// FindAllHandler ...
func (h *UserHandler) FindAllHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.FindAll()
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Create(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// UpdateHandler ...
func (h *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	req := request.UserUpdateRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Update(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// DeleteHandler ...
func (h *UserHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	req := request.UserDeleteRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Delete(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}
