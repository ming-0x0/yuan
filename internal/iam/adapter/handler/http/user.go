package http

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/internal/iam/application/user"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type updateProfileRequest struct {
	FullName string `json:"full_name"`
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	var req updateProfileRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userService.UpdateProfile(c.Request().Context(), req.FullName); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "updated"})
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	acc, err := h.userService.GetProfile(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id":         acc.ID.String(),
		"email":      acc.Email,
		"full_name":  acc.FullName,
	})
}

func getUserIDFromContext(c echo.Context) (id.ID, error) {
	// Assuming Echo JWT middleware puts token in context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	idInt, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, err
	}

	return id.FromInt64(idInt), nil
}
