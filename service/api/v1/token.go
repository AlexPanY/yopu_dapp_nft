package v1

import (
	"errors"
	"math/big"

	"ypt_server/api"
	"ypt_server/api/v1/request"
	"ypt_server/errno"
	"ypt_server/pkg/ether"
	"ypt_server/token"

	"github.com/gin-gonic/gin"
)

//CreateToken
func CreateToken(c *gin.Context) {
	var req request.CreateTokenRequest
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
	// api.SuccJSONWithData(c, 1)
	// return
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

//TransferToken
func TransferToken(c *gin.Context) {
	var req request.TransferTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	account, err := token.FindAccountByAddress(req.From)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	if account.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	t := token.ERC721_YopuNFT{
		Account: ether.NewAccount(1, req.From, req.Privatekey),
		TokenID: big.NewInt(req.TokenID),
	}

	if err := t.Get(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, 1)
}

//DescribeToken
func DescribeToken(c *gin.Context) {
	var req request.DescribeTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
}

//DescribeTokenList
func DescribeTokenList(c *gin.Context) {
	var req request.DescribeTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	// tokens := token.FindTokensByAccountID(req.AccountID)

	api.SuccJSONWithData(c, 1)
}
