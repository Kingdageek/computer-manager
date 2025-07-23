package controllers

import (
	"computer-manager/internal/api"
	"computer-manager/internal/services"

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
	reqCtx := ginCtx.Request.Context()
	data, err := c.svc.GetAllComputers(reqCtx)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}

func (c *ComputerController) GetByID(ginCtx *gin.Context) {

}

func (c *ComputerController) Create(ginCtx *gin.Context) {}

func (c *ComputerController) Update(ginCtx *gin.Context) {}

func (c *ComputerController) Delete(ginCtx *gin.Context) {}

func (c *ComputerController) AssignEmployee(ginCtx *gin.Context) {}
