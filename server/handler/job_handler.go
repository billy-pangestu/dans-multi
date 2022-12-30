package handler

import (
	"net/http"
	"strconv"

	"be-user-scheme/usecase"

	"github.com/go-chi/chi"
)

// JobHandler ...
type JobHandler struct {
	Handler
}

// FindAllHandler ...
func (h *JobHandler) FindAllHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		SendBadRequest(w, "Invalid page value")
		return
	}

	description := r.URL.Query().Get("description")
	location := r.URL.Query().Get("location")

	fulltime := r.URL.Query().Get("full_time")
	_, err = strconv.ParseBool(fulltime)
	if err != nil {
		fulltime = ""
	}

	jobUc := usecase.JobUC{ContractUC: h.ContractUC}
	res, err := jobUc.FindAll(page, location, description, fulltime)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}

// FindByIDHandlers ...
func (h *JobHandler) FindByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	id := chi.URLParam(r, "id")
	if id == "" {
		SendBadRequest(w, "Invalid parameter")
		return
	}

	jobUc := usecase.JobUC{ContractUC: h.ContractUC}
	res, err := jobUc.FindByID(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
}
