package v1

import (
	"errors"

	"ypt_server/api"
	"ypt_server/api/v1/request"
	"ypt_server/errno"
	"ypt_server/token"

	"github.com/gin-gonic/gin"
)

//CreateAccount
func CreateAccount(c *gin.Context) {
	var req request.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	if len(req.Nickname) <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("`nickname` must be required."))
		return
	}

	if len(req.Address) <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("`address` must be required."))
		return
	}

	a := &token.YPTAccount{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Address:  req.Address,
	}

	if err := a.Create(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	api.SuccJSONWithData(c, []int{})
}

//GetAccountByID - Get account info by account_id
func GetAccountByID(c *gin.Context) {
	var req request.GetAccountByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if req.AccountID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("`account_id` must be required."))
		return
	}

	account, err := token.FindAccountByID(req.AccountID)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, account)
	return
}

//DescribeAccountAssets Get account assets
func DescribeAccountAssets(c *gin.Context) {
	var req request.GetAccountByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}
	if req.AccountID <= 0 {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, errors.New("`account_id` must be required."))
		return
	}

	a, err := token.FindAccountByID(req.AccountID)
	if err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	if err := (&a).BalanceOf(); err != nil {
		api.ErrJSONWithRawErr(c, errno.ErrParamInvalid, err)
		return
	}

	api.SuccJSONWithData(c, a)
	return
}
