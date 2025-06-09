package handlers

import (
	"net/http"

	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
)

func NewParameterHandler(paramService service.ParameterService) ParameterHandler {
	return &paramHandler{paramService: paramService}
}

type paramHandler struct {
	paramService service.ParameterService
}

// GetAllParameters godoc
//
//	@Summary		Get all the parameters
//	@Description	Retrieve a list of all parameters
//	@Tags			parameters
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Parameter
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/api/params [get]
func (p *paramHandler) GetAllParameters(c *gin.Context) {
	params, err := p.paramService.GetAllParams()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, params)
}

// UpdateParameters godoc
//
//	@Summary		Update parameters
//	@Description	Update all parameters
//	@Tags			parameters
//	@Accept			json
//	@Produce		json
//	@Param			param	body		models.Parameter	true	"Parameter information"
//	@Success		200		{array}		models.Parameter
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/params [put]
func (p *paramHandler) UpdateParameters(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot find validated model in context", nil))
		return
	}
	paramRequest, ok := validatedModel.(*models.Parameter)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot cast validated model to ParameterRequest", nil))
		return
	}
	paramResponse, err := p.paramService.UpdateParams(paramRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, paramResponse)
}
