package request

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
}

//GetAccountByIDRequest
type GetAccountByIDRequest struct {
	AccountID int64 `json:"account_id"`
}
