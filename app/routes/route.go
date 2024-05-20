package routes

import (
	"myTaskApp/app/middlewares"
	_userData "myTaskApp/features/user/data"
	_userHandler "myTaskApp/features/user/handler"
	_userService "myTaskApp/features/user/service"
	encrypts "myTaskApp/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	dataService := _userData.New(db)
	userService := _userService.New(dataService, hashService)
	userHandlerAPI := _userHandler.New(userService)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())

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
