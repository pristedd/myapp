package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"myapp/handler"
	logic "myapp/logic"
	"myapp/repository"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	_, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}
	e := echo.New()

	repos := repository.PersonRepository{}
	logics := logic.PersonLogic{repos}
	handlers := handler.PersonHandler{logics}

	e.GET("/person", handlers.GetPerson)
	e.GET("/person/:id", handlers.GetPerson)
	e.POST("/person", handlers.CreatePerson)
	e.PUT("/person/:id", handlers.UpdatePerson)
	e.DELETE("/person/:id", handlers.DeletePerson)

	logrus.Fatal(e.Start(":8000"))
}
