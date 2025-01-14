package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodrigoapolo/api-students/db"
)

type API struct {
	Echo *echo.Echo
	DB *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dataBase := db.Init()
	studentDB := db.NewStudentHanler(dataBase)

	return &API{
		Echo: e,
		DB: studentDB,
	}
}

func (api *API) Start() error {
	err := api.Echo.Start(":8080");
	if  err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
		return fmt.Errorf("failed to start server: %w", err)
	}

	return err
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudents)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}