package handlers

import (
	"strconv"

	"hexagonal-gofr/internal/core/domain"
	"hexagonal-gofr/internal/core/ports"

	"gofr.dev/pkg/gofr"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx *gofr.Context) (interface{}, error) {
	var user domain.User
	if err := ctx.Bind(&user); err != nil {
		return nil, err
	}

	if err := h.service.CreateUser(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserHandler) GetUserByID(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
