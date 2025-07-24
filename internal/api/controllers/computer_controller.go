package controllers

import (
	"computer-manager/internal/api"
	"computer-manager/internal/api/http_errors"
	"computer-manager/internal/api/requests"
	"computer-manager/internal/dtos"
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
	reqCtx := ginCtx.Request.Context()
	var req requests.GetAllComputersRequest
	employeeCodes := ginCtx.QueryArray("employee_codes")
	req.EmployeeCodes = employeeCodes
	data, err := c.svc.GetAllComputers(reqCtx, &req)
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

func (c *ComputerController) validateComputerDto(dto *dtos.ComputerDto) error {
	if dto.Name == "" || dto.MacAddress == "" || dto.IPAddress == "" {
		return http_errors.NewBadRequestError("Name, MacAddress, and IPAddress are required")
	}
	if dto.EmployeeCode != nil && (*dto.EmployeeCode == "" || len(*dto.EmployeeCode) != 3) {
		return http_errors.NewBadRequestError("Employee code must be 3 characters long")
	}
	return nil
}

func (c *ComputerController) Create(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	var computerDto dtos.ComputerDto
	if err := ginCtx.ShouldBindJSON(&computerDto); err != nil {
		api.ErrorResponse(ginCtx, http_errors.NewBadRequestError("Invalid request body"))
		return
	}

	err := c.validateComputerDto(&computerDto)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}

	data, err := c.svc.CreateComputer(ctx, &computerDto)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}

func (c *ComputerController) Update(ginCtx *gin.Context) {
	computerIdStr := ginCtx.Param("id")
	computerId, err := strconv.ParseUint(computerIdStr, 10, 0)
	if err != nil {
		api.ErrorResponse(ginCtx, http_errors.NewBadRequestError("Invalid computer ID"))
		return
	}
	ctx := ginCtx.Request.Context()
	var computerDto dtos.ComputerDto
	if err = ginCtx.ShouldBindJSON(&computerDto); err != nil {
		api.ErrorResponse(ginCtx, http_errors.NewBadRequestError("Invalid request body"))
		return
	}

	err = c.validateComputerDto(&computerDto)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}

	data, err := c.svc.UpdateComputer(ctx, uint(computerId), &computerDto)
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}

func (c *ComputerController) Delete(ginCtx *gin.Context) {
	reqCtx := ginCtx.Request.Context()
	computerIdStr := ginCtx.Param("id")
	computerId, err := strconv.ParseUint(computerIdStr, 10, 0)
	if err != nil {
		api.ErrorResponse(ginCtx, http_errors.NewBadRequestError("Invalid computer ID"))
		return
	}
	data, err := c.svc.DeleteComputer(reqCtx, uint(computerId))
	if err != nil {
		api.ErrorResponse(ginCtx, err)
		return
	}
	api.SuccessResponse(ginCtx, data)
}
