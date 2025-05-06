package service

import (
	"net/http"
	"project_structure/internal/core/model"
	"project_structure/internal/core/ports"
	"project_structure/internal/utils"
)

type authService struct {
	resp ports.IResponse
}

func NewAuthServiceObj(resp ports.IResponse) ports.IAuthService {
	return &authService{
		resp: resp,
	}
}

func (a *authService) Register(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	a.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (a *authService) Login(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	a.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}
