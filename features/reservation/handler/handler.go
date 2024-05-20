package handler

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils/responses"
	"github.com/labstack/echo/v4"
	"net/http"
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
	requestCore := reservation.Core{
		UserID:     uint(idToken),
		HomestayID: newRequest.HomestayID,
		StartDate:  newRequest.StartDate,
		EndDate:    newRequest.EndDate,
	}

	errCreate := r.ReservationService.CreateReservation(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create reservation: "+errCreate.Error(), nil))
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

	requestCore := reservation.Core{
		HomestayID: newRequest.HomestayID,
		StartDate:  newRequest.StartDate,
		EndDate:    newRequest.EndDate,
	}

	errCheck := r.ReservationService.CheckAvailability(requestCore)
	if errCheck != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error check availability: "+errCheck.Error(), nil))
	}

	var response = AvailableResponse{
		Status: "available",
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success check availability", response))
}
