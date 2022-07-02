package ether

import (
	"ypt_server/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	defaultRawURL = "http://127.0.0.1:7545"
)

//NewEtherClientConn
func NewEtherClientConn() (*ethclient.Client, error) {
	rawURL := defaultRawURL
	if len(config.G.Contract.Host) <= 0 {
		rawURL = config.G.Contract.Host
	}
	client, err := ethclient.Dial(rawURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}
