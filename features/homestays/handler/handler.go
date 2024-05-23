package handler

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils/upload"
	"net/http"
	"strconv"

	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils/responses"

	"github.com/labstack/echo/v4"
)

type HomestayHandler struct {
	homestayService homestay.ServiceInterface
	homestayData    homestay.DataInterface
}

func New(hh homestay.ServiceInterface, hd homestay.DataInterface) *HomestayHandler {
	return &HomestayHandler{
		homestayService: hh,
		homestayData:    hd,
	}
}

func (h *HomestayHandler) CreateHomestay(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := HomestayRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	imageFiles := []string{"images1", "images2", "images3"}
	uploadUrls := make([]string, len(imageFiles))

	for i, image := range imageFiles {
		formHeader, err := c.FormFile(image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error formheader: "+err.Error(), nil))
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error formfile: "+err.Error(), nil))
		}
		defer formFile.Close()

		uploadUrl, err := upload.ImageUploadHelper(formFile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error upload: "+err.Error(), nil))
		}

		uploadUrls[i] = uploadUrl
	}

	requestCore := homestay.Core{
		UserID:       uint(idToken),
		HomestayName: newRequest.HomestayName,
		Description:  newRequest.Description,
		Address:      newRequest.Address,
		Images1:      uploadUrls[0],
		Images2:      uploadUrls[1],
		Images3:      uploadUrls[2],
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
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all homestay: "+err.Error(), nil))
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
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all homestay: "+errGetAll.Error(), nil))
	}

	var allHomestayResponse []HomestayResponse
	for _, v := range result {
		allHomestayResponse = append(allHomestayResponse, HomestayResponse{
			ID:            v.ID,
			HomestayName:  v.HomestayName,
			Address:       v.Address,
			PricePerNight: v.CostPerNight,
		})
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all homestay", allHomestayResponse))
}

func (h *HomestayHandler) GetMyHomestay(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.homestayService.GetMyHomestay(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get your homestay: "+err.Error(), nil))
	}

	var allMyHomestay []HomestayResponse
	for _, v := range result {
		allMyHomestay = append(allMyHomestay, HomestayResponse{
			ID:            v.ID,
			HomestayName:  v.HomestayName,
			Address:       v.Address,
			PricePerNight: v.CostPerNight,
			Description:   v.Description,
			Images1:       v.Images1,
			Images2:       v.Images2,
			Images3:       v.Images3,
		})
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all your homestay", allMyHomestay))
}

func (h *HomestayHandler) GetHomestayById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id:"+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.homestayService.GetHomestayById(uint(idConv), uint(idToken))
	var emptyHomestay homestay.Core
	if err != nil {
		if result == emptyHomestay {
			return c.JSON(http.StatusNotFound, responses.WebJSONResponse("error get homestay: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get homestay: "+err.Error(), nil))
	}

	responseResult := HomestayResponseById{
		ID:            uint(idConv),
		HomestayName:  result.HomestayName,
		Address:       result.Address,
		PricePerNight: result.CostPerNight,
		Description:   result.Description,
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

	imageFiles := []string{"images1", "images2", "images3"}
	uploadUrls := make([]string, len(imageFiles))

	for i, image := range imageFiles {
		formHeader, err := c.FormFile(image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error : ", "You need to fill images1-images2-images3 field"))
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error formfile: "+err.Error(), nil))
		}
		defer formFile.Close()

		uploadUrl, err := upload.ImageUploadHelper(formFile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error upload: "+err.Error(), nil))
		}

		uploadUrls[i] = uploadUrl
	}

	updateCore := homestay.Core{
		HomestayName: updateRequest.HomestayName,
		Description:  updateRequest.Description,
		Address:      updateRequest.Address,
		Images1:      uploadUrls[0],
		Images2:      uploadUrls[1],
		Images3:      uploadUrls[2],
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

	result, err := h.homestayData.GetHomestayByUserId(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update homestay: "+err.Error(), nil))
	}
	if len(result) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete homestay. if you want add homestay, please make a host again", nil))

	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete homestay", nil))

}

func (h *HomestayHandler) MakeHost(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := HomestayRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	imageFiles := []string{"images1", "images2", "images3"}
	uploadUrls := make([]string, len(imageFiles))

	for i, image := range imageFiles {
		formHeader, err := c.FormFile(image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error : ", "You need to fill images1-images2-images3 field"))
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error formfile: "+err.Error(), nil))
		}
		defer formFile.Close()

		uploadUrl, err := upload.ImageUploadHelper(formFile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error upload: "+err.Error(), nil))
		}

		uploadUrls[i] = uploadUrl
	}

	requestCore := homestay.Core{
		UserID:       uint(idToken),
		HomestayName: newRequest.HomestayName,
		Description:  newRequest.Description,
		Address:      newRequest.Address,
		Images1:      uploadUrls[0],
		Images2:      uploadUrls[1],
		Images3:      uploadUrls[2],
		CostPerNight: newRequest.PricePerNight,
	}

	err := h.homestayService.MakeHost(uint(idToken), requestCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error make a host: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success add homestay. Congratulations! you're host now", nil))
}
