package api

import (
	"computer-manager/internal/api/http_errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, err error) {
	var apiErr *http_errors.HttpError
	if ok := errors.As(err, &apiErr); ok {
		c.JSON(apiErr.HTTPCode, Response{
			Success: false,
			Error:   apiErr,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   err.Error(),
	})
}
