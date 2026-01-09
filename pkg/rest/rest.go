package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ming-0x0/yuan/internal/common/apperror"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Unknown error",
			Code:    apperror.Unknown.String(),
		})
		return
	}

	var appErr *apperror.AppError
	if apperror.As(err, &appErr) {
		status := mapAppErrorToHTTPStatus(appErr.Code())
		c.JSON(status, Response{
			Success: false,
			Message: appErr.Message(),
			Code:    appErr.Code().String(),
		})
		return
	}

	// Default for other errors (like binding errors)
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: err.Error(),
		Code:    apperror.InvalidArgument.String(),
	})
}

func mapAppErrorToHTTPStatus(code apperror.ErrorCode) int {
	switch code {
	case apperror.OK:
		return http.StatusOK
	case apperror.InvalidArgument:
		return http.StatusBadRequest
	case apperror.Unauthenticated:
		return http.StatusUnauthorized
	case apperror.PermissionDenied:
		return http.StatusForbidden
	case apperror.NotFound:
		return http.StatusNotFound
	case apperror.AlreadyExists:
		return http.StatusConflict
	case apperror.Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
