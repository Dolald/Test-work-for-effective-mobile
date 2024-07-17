package app

import (
	"os"
	testwork "test"
	handler "test/internal/handler/http"
	"test/internal/repository"
	"test/internal/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logrus.Info("set godotenv")

	db, err := repository.NewPostgreDB(repository.Configs{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		panic(err)
	}

	// ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	// defer cancel()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := new(testwork.Server) // создаём новый сервер

	if err := server.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}
