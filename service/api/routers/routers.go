package routers

import (
	m "ypt_server/api/routers/middleware"
	v1 "ypt_server/api/v1"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	// Middleware - http
	g.Use(gin.Recovery())
	g.Use(m.NoCache)
	g.Use(m.Options)
	g.Use(m.Secure)
	g.Use(m.RequestID())

	v1Api := g.Group("/v1")
	{
		v1Api.POST("/account/get", v1.GetAccountByID)
		v1Api.POST("/account/create", v1.CreateAccount)
		v1Api.POST("/account/assets", v1.DescribeAccountAssets)

		//nft token
		v1Api.POST("/token/create", v1.CreateToken)
		v1Api.POST("/token/list", v1.DescribeTokenList)
		v1Api.POST("/token/get", v1.DescribeToken)
		v1Api.POST("/token/buy", v1.TransferToken)
	}

	return g
}
