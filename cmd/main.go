package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	answerUsecase "github.com/mahauni/euro-farma-api/internal/answers/usecase"
	"github.com/mahauni/euro-farma-api/internal/handlers"
	"github.com/mahauni/euro-farma-api/internal/infra/repository"
	questionUsecase "github.com/mahauni/euro-farma-api/internal/questions/usecase"
	quizUsecase "github.com/mahauni/euro-farma-api/internal/quizzes/usecase"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load the .env: %v\n", err)
		os.Exit(1)
	}

	// dbConnStr := database.DatabaseConnectionString{
	// 	PostgresUser:     os.Getenv("POSTGRES_USER"),
	// 	PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	// 	PostgresHost:     os.Getenv("POSTGRES_HOST"),
	// 	PostgresPort:     os.Getenv("POSTGRES_PORT"),
	// 	PostgresDb:       os.Getenv("POSTGRES_DB"),
	// }
	// connString := database.CreateConnectionString(dbConnStr)
	// conn, err := pgx.Connect(context.Background(), connString)

	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL")),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	quizRepo := repository.NewQuizRepositoryPostgres(conn)
	questionRepo := repository.NewQuestionRepositoryPostgres(conn)
	answerRepo := repository.NewAnswerRepositoryPostgres(conn)

	quizUC := quizUsecase.NewCreateQuizUseCase(quizRepo)
	questionUC := questionUsecase.NewCreateQuestionUseCase(questionRepo)
	answerUC := answerUsecase.NewCreateAnswerUseCase(answerRepo)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})

		handlers.NewQuizHandler(r, quizUC)
		handlers.NewQuestionHandler(r, questionUC)
		handlers.NewAnswerHandler(r, answerUC)
	})

	http.ListenAndServe(":3333", r)
}
