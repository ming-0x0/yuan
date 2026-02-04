package http

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ming-0x0/yuan/internal/blog/application"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type BlogHandler struct {
	service application.BlogService
}

func NewBlogHandler(service application.BlogService) *BlogHandler {
	return &BlogHandler{
		service: service,
	}
}

type createRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *BlogHandler) Create(c echo.Context) error {
	var req createRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	authorID, err := getUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	if err := h.service.Create(c.Request().Context(), req.Title, req.Content, authorID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "created"})
}

type updateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *BlogHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var req updateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	authorID, err := getUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	if err := h.service.Update(c.Request().Context(), id.FromInt64(idInt), req.Title, req.Content, authorID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "updated"})
}

func (h *BlogHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	authorID, err := getUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	if err := h.service.Delete(c.Request().Context(), id.FromInt64(idInt), authorID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}

func (h *BlogHandler) Get(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	blog, err := h.service.Get(c.Request().Context(), id.FromInt64(idInt))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	blogs, err := h.service.List(c.Request().Context(), page, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, blogs)
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
