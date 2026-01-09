package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ming-0x0/yuan/internal/account/domain/account"
	"github.com/ming-0x0/yuan/pkg/rest"
)

type AccountHandler struct {
	accountApp account.AccountApp
}

func NewHandler(accountApp account.AccountApp) *AccountHandler {
	return &AccountHandler{accountApp: accountApp}
}

type CreateAccountRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rest.Error(c, err)
		return
	}

	err := h.accountApp.CreateAccount(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		rest.Error(c, err)
		return
	}

	rest.OK(c, nil)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AccountHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rest.Error(c, err)
		return
	}

	acc, err := h.accountApp.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		rest.Error(c, err)
		return
	}

	rest.OK(c, gin.H{
		"id":    acc.ID(),
		"email": acc.Email(),
	})
}

func (h *AccountHandler) GetAccountByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		rest.Error(c, nil) // Should handle missing param error
		return
	}

	acc, err := h.accountApp.GetAccountByEmail(c.Request.Context(), email)
	if err != nil {
		rest.Error(c, err)
		return
	}

	rest.OK(c, gin.H{
		"id":    acc.ID(),
		"email": acc.Email(),
	})
}
