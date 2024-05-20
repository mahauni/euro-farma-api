package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahauni/euro-farma-api/internal/users/entity"
	"github.com/mahauni/euro-farma-api/internal/users/usecase"
)

type userHandler struct {
	UserUsecase *usecase.UserUsecase
}

type userContext struct{}

func NewUserHandler(r chi.Router, useCase *usecase.UserUsecase) {
	handler := &userHandler{
		UserUsecase: useCase,
	}

	r.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetAllUserzes)
		r.Post("/", handler.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.UserCtx)
			r.Get("/", handler.GetUser)
			r.Put("/", handler.UpdateUser)
			r.Delete("/", handler.DeleteUser)
		})
	})
}

func (h *userHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		user, err := h.UserUsecase.FindUserById(r.Context(), userId)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), userContext{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *userHandler) GetAllUserzes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.UserUsecase.FindAllUsers(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", users)))
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, ok := ctx.Value(userContext{}).(*entity.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", user)))
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, ok := ctx.Value(userContext{}).(*entity.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = h.UserUsecase.UpdateUser(ctx, &user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", user)))
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, ok := ctx.Value(userContext{}).(*entity.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err := h.UserUsecase.DeleteUser(ctx, user.Id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", user)))
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.UserUsecase.CreateUser(ctx, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", user)))
}
