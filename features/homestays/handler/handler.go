package handler

import (
	"myTaskApp/app/middlewares"
	homestay "myTaskApp/features/homestays"
	"myTaskApp/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HomestayHandler struct {
	homestayService homestay.ServiceInterface
}

func New(hh homestay.ServiceInterface) *HomestayHandler {
	return &HomestayHandler{
		homestayService: hh,
	}
}

func (h *HomestayHandler) CreateHomestay(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := HomestayRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	requestCore := homestay.Core{
		UserID:       uint(idToken),
		HomestayName: newRequest.HomestayName,
		Description:  newRequest.Description,
		Address:      newRequest.Address,
		Images1:      newRequest.Images1,
		Images2:      newRequest.Images2,
		Images3:      newRequest.Images3,
		CostPerNight: newRequest.PricePerNight,
	}

	errCreate := h.homestayService.Create(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create homestay: "+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create homestay", nil))
}

func (h *HomestayHandler) GetAllForUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.homestayService.GetAllForUser(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all homestay "+err.Error(), nil))
	}

	var allForUserResponse []HomestayResponse
	for _, v := range result {
		allForUserResponse = append(allForUserResponse, HomestayResponse{
			ID:            v.ID,
			HomestayName:  v.HomestayName,
			Address:       v.Address,
			PricePerNight: v.CostPerNight,
		})
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all homestay", allForUserResponse))
}

func (h *HomestayHandler) GetAllHomestay(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, errGetAll := h.homestayService.GetAll(uint(idToken))
	if errGetAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all data "+errGetAll.Error(), nil))
	}

	var allHomestayResponse []HomestayResponse
	for _, v := range result {
		allHomestayResponse = append(allHomestayResponse, HomestayResponse{
			ID:           v.ID,
			HomestayName: v.HomestayName,
		})
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all homestay", allHomestayResponse))
}

func (h *HomestayHandler) GetHomestayById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id:"+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.homestayService.GetHomestayById(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get homestay "+err.Error(), nil))
	}

	responseResult := HomestayResponseById{
		ID:           uint(idConv),
		HomestayName: result.HomestayName,
		Description:  result.Description,
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get homestay", responseResult))
}

func (h *HomestayHandler) UpdateHomestay(c echo.Context) error {
	id := c.Param("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id: "+errConv.Error(), nil))
	}

	updateRequest := HomestayRequest{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}
	updateCore := homestay.Core{
		HomestayName: updateRequest.HomestayName,
		Description:  updateRequest.Description,
		Address:      updateRequest.Address,
		Images1:      updateRequest.Images1,
		Images2:      updateRequest.Images2,
		Images3:      updateRequest.Images3,
		CostPerNight: updateRequest.PricePerNight,
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.homestayService.Update(uint(idInt), uint(idToken), updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update homestay: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success update homestay", nil))
}

func (h *HomestayHandler) DeleteHomestay(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id: "+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	tx := h.homestayService.Delete(uint(idConv), uint(idToken))
	if tx != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error delete homestay: "+tx.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete homestay", nil))
}
