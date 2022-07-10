package v1

import (
	"errors"
	"math/big"
	"strconv"

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
	a, err := token.FindAccountByAddress(t.Account.Address)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if a.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	t.Account.AccountID = uint64(a.ID)

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

	a, err := token.FindAccountByAddress(req.BuyerAddress)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	if a.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	a.Privatekey = req.BuyerPrivatekey

	if err := a.Buy(req.TokenID); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, 1)
}

//SetTokenPrice
func SetTokenPrice(c *gin.Context) {
	var req request.SetTokenPriceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	t := token.ERC721_YopuNFT{
		TokenID: big.NewInt(req.TokenID),
		Account: ether.NewAccount(1, req.Address, req.Privatekey),
	}

	a, err := token.FindAccountByAddress(t.Account.Address)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if a.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	price, err := strconv.ParseInt(req.Price, 10, 64)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if err := t.SetPrice(price); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
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

	t := token.ERC721_YopuNFT{
		Account: ether.NewAccount(1, req.Address, req.Privatekey),
		TokenID: big.NewInt(req.TokenID),
	}

	a, err := token.FindAccountByAddress(t.Account.Address)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if a.ID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("account is not exists"))
		return
	}

	count, err := (&a).BalanceOf()
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if t.TokenID.Cmp(count) >= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("`TokenID` is not exists"))
		return
	}

	if err := t.Get(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, t)
}

//DescribeTokenList
func DescribeTokenList(c *gin.Context) {
	var req request.DescribeTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	tokens := token.FindAllTokens()
	if len(tokens) <= 0 {
		api.SuccJSONWithData(c, tokens)
		return
	}

	yptTokens := make([]*token.ERC721_YopuNFT, 0, len(tokens))
	for _, v := range tokens {
		v := v
		t := &token.ERC721_YopuNFT{
			TokenID: big.NewInt(v.ID),
		}
		if err := t.Get(); err != nil {
			continue
		}
		yptTokens = append(yptTokens, t)
	}

	api.SuccJSONWithData(c, yptTokens)
}
