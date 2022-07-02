package token

import (
	"context"
	"fmt"

	"ypt_server/config"
	"ypt_server/contracts"
	"ypt_server/pkg/ether"
	"ypt_server/pkg/ifps"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	defaultContractAddress = "0x0Ad6C6B9cAd650d5A61D883cc7F2D15B49697028"
)

//ERC721_YopuNFTProperty
type ERC721_YopuNFTProperty struct {
	field string `json:"field"`
	Value string `json:"value"`
}

//ERC721_YopuNFT
type ERC721_YopuNFT struct {
	Account     ether.Account             `json:"-"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Collection  string                    `json:"collection"`
	Properties  []*ERC721_YopuNFTProperty `json:"properties"`
}

//Mint Minting a single token based on ERC721.
func (t *ERC721_YopuNFT) Mint() error {
	cid, err := t.NewMetadataURI()
	if err != nil {
		return err
	}
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return err
	}

	contractAddress := defaultContractAddress
	if len(config.G.Contract.Address) <= 0 {
		contractAddress = config.G.Contract.Address
	}

	ypt, err := contracts.NewYopuNFT(ether.HexToAddress(contractAddress), conn)
	if err != nil {
		return err
	}

	chainID, err := conn.NetworkID(context.Background())
	if err != nil {
		return err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.Account.ECDSA_PrivateKey, chainID)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://ipfs.io/%s", cid)
	transaction, err := ypt.SafeMint(transactOpts, t.Account.EthAddress, uri)
	if err != nil {
		return err
	}

	fmt.Println(transaction)
	return nil
}

//NewMetadataURI
func (t *ERC721_YopuNFT) NewMetadataURI() (string, error) {
	return ifps.UpdateMetadataWithJSON(t)
}
