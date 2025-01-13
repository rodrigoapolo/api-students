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

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
	return c.String(http.StatusCreated, "Create students")
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStud)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, updateStud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("DELETE %s student", id)
	return c.String(http.StatusOK, deleteStud)
}
