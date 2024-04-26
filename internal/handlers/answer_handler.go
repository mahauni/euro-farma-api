package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahauni/euro-farma-api/internal/answers/entity"
	"github.com/mahauni/euro-farma-api/internal/answers/usecase"
)

type answerHandler struct {
	AnswerUsecase *usecase.AnswerUsecase
}

type answerContext struct{}

func NewAnswerHandler(r chi.Router, useCase *usecase.AnswerUsecase) {
	handler := &answerHandler{
		AnswerUsecase: useCase,
	}

	r.Route("/answer", func(r chi.Router) {
		r.Get("/", handler.GetAllAnswers)
		r.Get("/quiz/{id}", handler.GetAllAnswersByQuiz)
		r.Get("/question/{id}", handler.GetAllAnswersByQuestion)
		r.Post("/", handler.CreateAnswer)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.AnswerCtx)
			r.Get("/", handler.GetAnswer)
			r.Put("/", handler.UpdateAnswer)
			r.Delete("/", handler.DeleteAnswer)
		})
	})
}

func (h *answerHandler) AnswerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		answerID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		answer, err := h.AnswerUsecase.FindAnswerById(r.Context(), answerID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), answerContext{}, answer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *answerHandler) GetAllAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	answers, err := h.AnswerUsecase.FindAllAnswers(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answers)))
}

func (h *answerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	answer, ok := ctx.Value(answerContext{}).(*entity.Answer)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answer)))
}

func (h *answerHandler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var answer entity.Answer
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, ok := ctx.Value(answerContext{}).(*entity.Answer)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = h.AnswerUsecase.UpdateAnswer(ctx, &answer)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answer)))
}

func (h *answerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	answer, ok := ctx.Value(answerContext{}).(*entity.Answer)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err := h.AnswerUsecase.DeleteAnswer(ctx, answer.ID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answer)))
}

func (h *answerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var answer entity.Answer
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.AnswerUsecase.CreateAnswer(ctx, &answer)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answer)))
}

func (h *answerHandler) GetAllAnswersByQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quizId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	answers, err := h.AnswerUsecase.FindAllAnswersByQuiz(ctx, quizId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answers)))
}

func (h *answerHandler) GetAllAnswersByQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	questionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	answers, err := h.AnswerUsecase.FindAllAnswersByQuestion(ctx, questionId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", answers)))
}
