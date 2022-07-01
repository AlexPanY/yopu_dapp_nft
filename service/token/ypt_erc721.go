package token

import (
	"fmt"

	"ypt_server/contracts"
	"ypt_server/pkg/ether"
	"ypt_server/pkg/ifps"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	defaultContractAddress = "0x546A1878AA552B5dC832Cf32cAeCBA6a87ffD975"
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

	ypt, err := contracts.NewYopuNFT(ether.HexToAddress(defaultContractAddress), conn)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://ipfs.io/%s", cid)
	transaction, err := ypt.SafeMint(t.Account.EthAddress, uri)
	if err != nil {
		return err
	}

	transactOpts, err := bind.NewTransactorWithChainID(key, "test", chainID)
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
