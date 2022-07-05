package v1

import (
	"errors"
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
	api.SuccJSONWithData(c, 1)
	return
	account, err := token.FindAccountByAddress(t.Account.Address)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if account.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	t.Account.AccountID = uint64(account.ID)

	if err := t.Mint(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, 1)
}

type DescribeAssetListRequest struct {
	AccountID int64 `json:"account_id"`
}

//DescribeAssetList
func DescribeAssetList(c *gin.Context) {
	var req DescribeAssetListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	tokens := token.FindTokensByAccountID(req.AccountID)

	api.SuccJSONWithData(c, tokens)
}
