package handler

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type FeedbackHandler struct {
	FeedbackService feedback.ServiceInterface
}

func NewFeedbackHandler(FeedbackService feedback.ServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		FeedbackService: FeedbackService,
	}
}

func (f *FeedbackHandler) CreateFeedback(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := FeedbackRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	requestCore := feedback.Core{
		UserID:     uint(idToken),
		HomestayID: newRequest.HomestayID,
		Rating:     newRequest.Rating,
		Feedback:   newRequest.Feedback,
	}

	errCreate := f.FeedbackService.CreateFeedback(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create feedback: "+errCreate.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success create feedback", nil))
}

func (f *FeedbackHandler) GetFeedbackByHomestayId(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id: "+errConv.Error(), nil))
	}

	result, errGet := f.FeedbackService.GetFeedbackByHomestayId(uint(idConv))
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all feedback: "+errGet.Error(), nil))
	}

	var allFeedbackResponse []FeedbackResponse
	for _, v := range result {
		allFeedbackResponse = append(allFeedbackResponse, FeedbackResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			HomestayID: v.HomestayID,
			Rating:     v.Rating,
			Feedback:   v.Feedback,
			CreatedAt:  v.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all feedback", allFeedbackResponse))
}
