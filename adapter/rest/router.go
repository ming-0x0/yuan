package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ming-0x0/yuan/adapter/rest/account"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(
	accountHandler *account.AccountHandler,
) *Router {
	engine := gin.Default()

	v1 := engine.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.POST("", accountHandler.CreateAccount)
			accounts.POST("/login", accountHandler.Login)
			accounts.GET("/:email", accountHandler.GetAccountByEmail)
		}
	}

	return &Router{engine: engine}
}

func (r *Router) Run(port int) error {
	return r.engine.Run(fmt.Sprintf(":%d", port))
}
