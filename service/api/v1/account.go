package v1

import (
	"errors"

	"ypt_server/api"
	"ypt_server/errno"
	"ypt_server/token"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
}

//CreateAccount
func CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
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

//GetAccountByIDRequest
type GetAccountByIDRequest struct {
	AccountID int64 `json:"account_id"`
}

//GetAccountByID - Get account info by account_id
func GetAccountByID(c *gin.Context) {
	var req GetAccountByIDRequest
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
