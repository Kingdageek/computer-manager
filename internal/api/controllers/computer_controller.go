package controllers

import (
	"computer-manager/internal/api"
	"computer-manager/internal/api/http_errors"
	"computer-manager/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComputerController struct {
	svc *services.ComputerService
}

func NewComputerController(svcs *services.Services) *ComputerController {
	return &ComputerController{
		svc: svcs.ComputerService,
	}
}

func (c *ComputerController) GetAll(ginCtx *gin.Context) {
	// so request can be cancelled if the client disconnects
	// or if the server is shutting down
	reqCtx := ginCtx.Request.Context()
	data, err := c.svc.GetAllComputers(reqCtx)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}

func (c *ComputerController) GetByID(ginCtx *gin.Context) {
	reqCtx := ginCtx.Request.Context()
	computerIdStr := ginCtx.Param("id")
	computerId, err := strconv.ParseUint(computerIdStr, 10, 0)
	if err != nil {
		api.ErrorResponse(ginCtx, http_errors.NewBadRequestError("Invalid computer ID"))
		return
	}
	data, err := c.svc.GetComputerByID(reqCtx, uint(computerId))
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}

func (c *ComputerController) Create(ginCtx *gin.Context) {}

func (c *ComputerController) Update(ginCtx *gin.Context) {}

func (c *ComputerController) Delete(ginCtx *gin.Context) {}

func (c *ComputerController) AssignEmployee(ginCtx *gin.Context) {}
