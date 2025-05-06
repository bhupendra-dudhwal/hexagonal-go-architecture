package service

import (
	"net/http"
	"project_structure/internal/core/model"
	"project_structure/internal/core/ports"
	"project_structure/internal/utils"
)

type userService struct {
	resp ports.IResponse
}

func NewUserServiceObj(resp ports.IResponse) ports.IUserService {
	return &userService{resp: resp}
}

func (u *userService) Add(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	u.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (u *userService) List(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	u.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (u *userService) Info(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	u.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (u *userService) Update(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	u.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}

func (u *userService) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetReqID(r.Context())
	u.resp.JSON(w, model.Response{
		StatusCode: true,
		RequestID:  reqID,
		Message:    "SUCCESS",
	}, http.StatusOK)
}
