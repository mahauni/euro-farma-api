package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mahauni/euro-farma-api/internal/questions/entity"
	"github.com/mahauni/euro-farma-api/internal/questions/usecase"
)

type questionHandler struct {
	QuestionUsecase *usecase.QuestionUsecase
}

type questionContext struct{}

func NewQuestionHandler(r chi.Router, useCase *usecase.QuestionUsecase) {
	handler := &questionHandler{
		QuestionUsecase: useCase,
	}

	r.Route("/question", func(r chi.Router) {
		r.Get("/", handler.GetAllQuestions)
		r.Get("/quiz/{id}", handler.GetAllQuestionsByQuiz)
		r.Post("/", handler.CreateQuestion)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.QuestionCtx)
			r.Get("/", handler.GetQuestion)
			r.Put("/", handler.UpdateQuestion)
			r.Delete("/", handler.DeleteQuestion)
		})
	})
}

func (h *questionHandler) QuestionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		questionID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		question, err := h.QuestionUsecase.FindQuestionById(r.Context(), questionID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), questionContext{}, question)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *questionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	questions, err := h.QuestionUsecase.FindAllQuestions(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", questions)))
}

func (h *questionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	question, ok := ctx.Value(questionContext{}).(*entity.Question)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", question)))
}

func (h *questionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var question entity.Question
	err := json.NewDecoder(r.Body).Decode(&question)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, ok := ctx.Value(questionContext{}).(*entity.Question)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err = h.QuestionUsecase.UpdateQuestion(ctx, &question)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", question)))
}

func (h *questionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	question, ok := ctx.Value(questionContext{}).(*entity.Question)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	err := h.QuestionUsecase.DeleteQuestion(ctx, question.ID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", question)))
}

func (h *questionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var question entity.Question
	err := json.NewDecoder(r.Body).Decode(&question)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.QuestionUsecase.CreateQuestion(ctx, &question)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", question)))
}

func (h *questionHandler) GetAllQuestionsByQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	quizId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	questions, err := h.QuestionUsecase.FindAllQuestionsByQuiz(ctx, quizId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", questions)))
}
