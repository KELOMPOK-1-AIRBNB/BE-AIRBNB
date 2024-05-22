package routes

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	_feedbackData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback/data"
	_feedbackHandler "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback/handler"
	_feedbackService "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback/service"
	_homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	_homestayHandler "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/handler"
	_homestayService "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/service"
	_reservationData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation/data"
	_reservationHandler "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation/handler"
	_reservationService "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation/service"
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

	homestayData := _homestayData.New(db, dataService)
	homestayService := _homestayService.New(homestayData, dataService)
	homestayHandler := _homestayHandler.New(homestayService, homestayData)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.POST("users/upgrade", homestayHandler.MakeHost, middlewares.JWTMiddleware())
	e.POST("/users/changeprofilepicture", userHandlerAPI.UpdateProfilePicture, middlewares.JWTMiddleware())

	e.GET("/homestays", homestayHandler.GetAllForUser, middlewares.JWTMiddleware())
	e.GET("/homestays/:id", homestayHandler.GetHomestayById, middlewares.JWTMiddleware())
	e.GET("/homestays/host", homestayHandler.GetAllHomestay, middlewares.JWTMiddleware())
	e.POST("/homestays/host", homestayHandler.CreateHomestay, middlewares.JWTMiddleware())
	e.GET("/homestays/host/myHomestay", homestayHandler.GetMyHomestay, middlewares.JWTMiddleware())
	e.PUT("/homestays/host/:id", homestayHandler.UpdateHomestay, middlewares.JWTMiddleware())
	e.DELETE("/homestays/host/:id", homestayHandler.DeleteHomestay, middlewares.JWTMiddleware())

	reservationData := _reservationData.New(db)
	reservationService := _reservationService.New(reservationData, dataService, homestayData)
	reservationHandler := _reservationHandler.NewReservationHandler(reservationService)

	e.POST("/reservations/availability", reservationHandler.CheckAvailability, middlewares.JWTMiddleware())
	e.POST("/reservations", reservationHandler.CreateReservation, middlewares.JWTMiddleware())
	e.GET("/reservations/history", reservationHandler.GetHistory, middlewares.JWTMiddleware())

	feedbackData := _feedbackData.New(db)
	feedbackService := _feedbackService.New(feedbackData, reservationData, homestayData)
	feedbackHandler := _feedbackHandler.NewFeedbackHandler(feedbackService)

	e.POST("/feedback", feedbackHandler.CreateFeedback, middlewares.JWTMiddleware())
	e.GET("/feedback/:id", feedbackHandler.GetFeedbackByHomestayId, middlewares.JWTMiddleware())

}

// contoh upload file
//	e.POST("/file", FileUpload)
//}
//
//
//
//func FileUpload(c echo.Context) error {
//	formHeader, err := c.FormFile("file")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]any{
//			"status": "gagal header",
//		})
//	}
//
//	formFile, err := formHeader.Open()
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]any{
//			"status": "gagal formfile",
//		})
//	}
//
//	uploadUrl, err := upload.ImageUploadHelper(formFile)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]any{
//			"status": "gagal upload",
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]any{
//		"status": "success",
//		"image":  uploadUrl,
//	})
//
//}

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
