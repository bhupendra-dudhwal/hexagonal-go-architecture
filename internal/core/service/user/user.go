package user

import (
	"net/http"
	"project_structure/internal/core/constants"
	"project_structure/internal/core/ports/response"
	"project_structure/internal/core/ports/user"
)

type userService struct {
	resp response.IResponse
}

func NewUserServiceObj(resp response.IResponse) user.IUser {
	return &userService{resp: resp}
}

func (u *userService) Add(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(constants.REQUEST_KEY).(string)
	u.resp.JSON(w, map[string]string{"status": "200", "msg": "add", "request_id": reqID}, http.StatusOK)
}

func (u *userService) GetAll(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(constants.REQUEST_KEY).(string)
	u.resp.JSON(w, map[string]string{"status": "200", "msg": "GetAll", "request_id": reqID}, http.StatusOK)
}

func (u *userService) Get(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(constants.REQUEST_KEY).(string)
	u.resp.JSON(w, map[string]string{"status": "200", "msg": "Get", "request_id": reqID}, http.StatusOK)
}

func (u *userService) Update(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(constants.REQUEST_KEY).(string)
	u.resp.JSON(w, map[string]string{"status": "200", "msg": "Update", "request_id": reqID}, http.StatusOK)
}

func (u *userService) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(constants.REQUEST_KEY).(string)
	u.resp.JSON(w, map[string]string{"status": "200", "msg": "Delete", "request_id": reqID}, http.StatusOK)
}
