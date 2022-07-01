package v1

import (
	"ypt_server/api"
	"ypt_server/errno"
	"ypt_server/pkg/ether"
	"ypt_server/token"

	"github.com/gin-gonic/gin"
)

//CreateYPTTokenRequest
type CreateYPTTokenRequest struct {
	Address     string                          `json:"address"`
	Privatekey  string                          `json:"privatekey"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Image       string                          `json:"image"`
	Collection  string                          `json:"collection"`
	Properties  []*token.ERC721_YopuNFTProperty `json:"properties"`
}

//CreateYPTToken
func CreateYPTToken(c *gin.Context) {
	var req CreateYPTTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	t := token.ERC721_YopuNFT{
		Account:     ether.NewAccount(1, req.Address, req.Privatekey),
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Collection:  req.Collection,
		Properties:  req.Properties,
	}

	if err := t.Mint(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, 1)
}

type DescribeAssetListRequest struct {
	Address string `json:"address"`
}

//DescribeAssetList
func DescribeAssetList(c *gin.Context) {
	var req DescribeAssetListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, 1)
}
