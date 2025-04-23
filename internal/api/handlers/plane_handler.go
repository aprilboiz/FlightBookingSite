package handlers

import (
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewPlaneHandler(planeService service.PlaneService) PlaneHandler {
	return &planeHandler{planeService: planeService}
}

type planeHandler struct {
	planeService service.PlaneService
}

// GetAllPlanes godoc
// @Summary Get all planes
// @Description Retrieve a list of all planes
// @Tags planes
// @Accept json
// @Produce json
// @Success 200 {array} dto.PlaneResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/planes [get]
func (h *planeHandler) GetAllPlanes(c *gin.Context) {
	planes, err := h.planeService.GetAllPlanes()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, planes)
}

// GetPlaneByCode godoc
// @Summary Get plane by code
// @Description Retrieve a plane by its unique code
// @Tags planes
// @Accept json
// @Produce json
// @Param code path string true "Plane Code"
// @Success 200 {object} dto.PlaneResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/planes/{code} [get]
func (h *planeHandler) GetPlaneByCode(c *gin.Context) {
	code := c.Param("code")
	plane, err := h.planeService.GetPlaneByCode(code)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, plane)
}
