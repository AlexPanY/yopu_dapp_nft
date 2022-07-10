package request

type CreateAccountRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
}

//GetAccountByIDRequest
type GetAccountByIDRequest struct {
	Address    string `json:"Address"`
	Privatekey string `json:"PrivateKey"`
	AccountID  int64  `json:"account_id"`
}

//DescribeAccountAssetsRequest
type DescribeAccountAssetsRequest struct {
	Address    string `json:"Address"`
	Privatekey string `json:"PrivateKey"`
}
