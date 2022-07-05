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

		//nft token
		v1Api.POST("/asset/create", v1.CreateYPTToken)
		v1Api.POST("/asset/list", v1.DescribeAssetList)
	}

	return g
}
