package service

import (
	"net/http"
	"project_structure/internal/core/model"
	"project_structure/internal/core/ports"
	"project_structure/internal/utils"
)

type healthService struct {
	resp ports.IResponse
}

func NewHealthServiceObj(resp ports.IResponse) ports.IHealthServices {
	return &healthService{resp: resp}
}

func (h *healthService) Readiness(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	h.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (h *healthService) Liveness(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	h.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}
