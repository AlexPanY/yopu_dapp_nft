package request

import (
	"ypt_server/token"
)

//CreateTokenRequest
type CreateTokenRequest struct {
	Address     string                          `json:"address"`
	Privatekey  string                          `json:"privatekey"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Image       string                          `json:"image"`
	Collection  string                          `json:"collection"`
	Properties  []*token.ERC721_YopuNFTProperty `json:"properties"`
}

//TransferTokenRequest
type TransferTokenRequest struct {
	From       string `json:"From"`
	Privatekey string `json:"Privatekey"`
	To         string `json:"To"`
	TokenID    int64  `json:"TokenID"`
}

type DescribeTokenRequest struct {
	TokenID int64 `json:"TokenID"`
}
