package request

import (
	"ypt_server/token"
)

//CreateTokenRequest
type CreateTokenRequest struct {
	Address     string                          `json:"Address"`
	Privatekey  string                          `json:"PrivateKey"`
	Name        string                          `json:"Name"`
	Description string                          `json:"Description"`
	Image       string                          `json:"Image"`
	Collection  string                          `json:"Collection"`
	Properties  []*token.ERC721_YopuNFTProperty `json:"Properties"`
	Price       float64                         `json:"Price"`
}

//TransferTokenRequest
type TransferTokenRequest struct {
	BuyerAddress    string `json:"BuyerAddress"`
	BuyerPrivatekey string `json:"BuyerPrivatekey"`
	TokenID         int64  `json:"TokenID"`
}

type DescribeTokenRequest struct {
	Address    string `json:"Address"`
	Privatekey string `json:"PrivateKey"`
	TokenID    int64  `json:"TokenID"`
}

//SetTokenPriceRequest
type SetTokenPriceRequest struct {
	Address    string `json:"Address"`
	Privatekey string `json:"PrivateKey"`
	TokenID    int64  `json:"TokenID"`
	Price      string `json:"Price"`
}
