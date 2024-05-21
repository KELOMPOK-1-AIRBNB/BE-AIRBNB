package routes

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	_homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	_homestayHandler "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/handler"
	_homestayService "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/service"
	_userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
	_userHandler "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/handler"
	_userService "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/service"
	encrypts "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	dataService := _userData.New(db)
	userService := _userService.New(dataService, hashService)
	userHandlerAPI := _userHandler.New(userService)

	homestayData := _homestayData.New(db)
	homestayService := _homestayService.New(homestayData, dataService)
	homestayHandler := _homestayHandler.New(homestayService)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())

	e.GET("/homestays", homestayHandler.GetAllForUser, middlewares.JWTMiddleware())
	e.GET("/homestays/host", homestayHandler.GetAllHomestay, middlewares.JWTMiddleware())
	e.POST("/homestays/host", homestayHandler.CreateHomestay, middlewares.JWTMiddleware())
	e.GET("/homestays/host/myHomestay", homestayHandler.GetMyHomestay, middlewares.JWTMiddleware())
	e.GET("/homestays/host/:id", homestayHandler.GetHomestayById, middlewares.JWTMiddleware())
	e.PUT("/homestays/host/:id", homestayHandler.UpdateHomestay, middlewares.JWTMiddleware())
	e.DELETE("/homestays/host/:id", homestayHandler.DeleteHomestay, middlewares.JWTMiddleware())

}

// projectData := _projectData.New(db)
// projectService := _projectService.New(projectData, dataService)
// projectHandlerAPI := _projectHandler.New(projectService)

// taskData := _taskData.New(db)
// taskService := _taskService.New(taskData, projectData)
// taskHandlerAPI := _taskHandler.New(taskService)

// e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
// e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
// e.GET("/projects/:id", projectHandlerAPI.GetProjectById, middlewares.JWTMiddleware())
// e.PUT("/projects/:id", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
// e.DELETE("/projects/:id", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

// e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
// e.PUT("/tasks/:id", taskHandlerAPI.UpdateTaskById, middlewares.JWTMiddleware())
// e.DELETE("/tasks/:id", taskHandlerAPI.DeleteTaskById, middlewares.JWTMiddleware())
