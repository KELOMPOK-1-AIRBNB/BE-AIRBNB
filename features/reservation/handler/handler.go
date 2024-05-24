package handler

import (
	"net/http"
	"time"

	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils/responses"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	ReservationService reservation.ServiceInterface
}

func NewReservationHandler(ReservationService reservation.ServiceInterface) *ReservationHandler {
	return &ReservationHandler{
		ReservationService: ReservationService,
	}
}

func (r *ReservationHandler) CreateReservation(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := ReservationRequest{}

	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	var dateLayout = "2006-01-02"
	startDate, _ := time.Parse(dateLayout, newRequest.StartDate)
	endDate, _ := time.Parse(dateLayout, newRequest.EndDate)

	requestCore := reservation.Core{
		UserID:     uint(idToken),
		HomestayID: newRequest.HomestayID,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	errCreate := r.ReservationService.CreateReservation(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create reservation: ", errCreate.Error()))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create reservation", nil))
}

func (r *ReservationHandler) GetHistory(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, errGetAll := r.ReservationService.GetHistory(uint(idToken))
	if errGetAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all data "+errGetAll.Error(), nil))
	}

	var allHistoryResponse []HistoryResponse
	for _, v := range result {
		allHistoryResponse = append(allHistoryResponse, HistoryResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			HomestayID: v.HomestayID,
			StartDate:  v.StartDate,
			EndDate:    v.EndDate,
			TotalPrice: v.TotalPrice,
			Status:     v.Status,
		})
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all history", allHistoryResponse))
}

func (r *ReservationHandler) CheckAvailability(c echo.Context) error {
	newRequest := AvailableRequest{}
	errBind := c.Bind(&newRequest)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	var dateLayout = "2006-01-02"
	startDate, errStart := time.Parse(dateLayout, newRequest.StartDate)
	endDate, errEnd := time.Parse(dateLayout, newRequest.EndDate)

	if errStart != nil || errEnd != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("invalid format date", nil))
	}

	now := time.Now().Truncate(24 * time.Hour)
	if startDate.Before(now) || endDate.Before(now) {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("date must be today or in the future", nil))
	}

	if startDate.After(endDate) {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("start date must be before end date", nil))
	}

	requestCore := reservation.Core{
		HomestayID: newRequest.HomestayID,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	errCheck := r.ReservationService.CheckAvailability(requestCore)
	if errCheck != nil {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success check availability: ", "not available"))
	}

	var response = AvailableResponse{
		Status: "available",
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success check availability", response))
}
