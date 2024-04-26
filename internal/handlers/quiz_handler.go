package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahauni/euro-farma-api/internal/quizzes/entity"
	"github.com/mahauni/euro-farma-api/internal/quizzes/usecase"
)

type quizHandler struct {
	QuizUsecase *usecase.QuizUsecase
}

type quizContext struct{}

func NewQuizHandler(r chi.Router, useCase *usecase.QuizUsecase) {
	handler := &quizHandler{
		QuizUsecase: useCase,
	}

	r.Route("/quiz", func(r chi.Router) {
		r.Get("/", handler.GetAllQuizzes)
		r.Post("/", handler.CreateQuiz)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.QuizCtx)
			r.Get("/", handler.GetQuiz)
			r.Put("/", handler.UpdateQuiz)
			r.Delete("/", handler.DeleteQuiz)
		})
	})
}

func (h *quizHandler) QuizCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quizID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		quiz, err := h.QuizUsecase.FindQuizById(r.Context(), quizID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), quizContext{}, quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *quizHandler) GetAllQuizzes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quizzes, err := h.QuizUsecase.FindAllQuiz(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", quizzes)))
}

func (h *quizHandler) GetQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quiz, ok := ctx.Value(quizContext{}).(*entity.Quiz)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", quiz)))
}

func (h *quizHandler) UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var quiz entity.Quiz
	err := json.NewDecoder(r.Body).Decode(&quiz)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, ok := ctx.Value(quizContext{}).(*entity.Quiz)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = h.QuizUsecase.UpdateQuiz(ctx, &quiz)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", quiz)))
}

func (h *quizHandler) DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quiz, ok := ctx.Value(quizContext{}).(*entity.Quiz)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err := h.QuizUsecase.DeleteQuiz(ctx, quiz.ID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", quiz)))
}

func (h *quizHandler) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var quiz entity.Quiz
	err := json.NewDecoder(r.Body).Decode(&quiz)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.QuizUsecase.CreateQuiz(ctx, &quiz)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", quiz)))
}
