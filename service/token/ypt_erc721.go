package token

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"ypt_server/config"
	"ypt_server/contracts"
	"ypt_server/pkg/ether"
	"ypt_server/pkg/ifps"
	"ypt_server/pkg/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	defaultContractAddress = "0x0Ad6C6B9cAd650d5A61D883cc7F2D15B49697028"
)

//ERC721_YopuNFTProperty
type ERC721_YopuNFTProperty struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

//ERC721_YopuNFT
type ERC721_YopuNFT struct {
	Account     ether.Account             `json:"-"`
	TokenID     *big.Int                  `json:"-"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Collection  string                    `json:"collection"`
	Properties  []*ERC721_YopuNFTProperty `json:"properties"`
}

//MetadataURI
func (t *ERC721_YopuNFT) MetadataURI() (string, error) {
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return "", err
	}

	contractAddress := defaultContractAddress
	if len(config.G.Contract.Address) > 0 {
		contractAddress = config.G.Contract.Address
	}

	caller, err := contracts.NewYopuNFTCaller(ether.HexToAddress(contractAddress), conn)
	if err != nil {
		return "", err
	}

	return caller.TokenURI(&bind.CallOpts{}, t.TokenID)
}

//Get
func (t *ERC721_YopuNFT) Get() error {
	uri, err := t.MetadataURI()
	if err != nil {
		return err
	}

	fmt.Println(uri)
	idx := strings.Index(uri, "https://ipfs.io/")
	if idx <= 0 {
		return errors.New("incorrect ifps URI format")
	}

	// testuri := "http://localhost:8080/ipfs/QmapAmUAS9Uou6xMTeZBtAgaVJhRskFJgL5RWmCfVSFe4M"
	resultByte, err := proxy.Get(testuri, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resultByte, &t); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%+v", *t))
	fmt.Println(string(resultByte))

	return nil
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
	if len(config.G.Contract.Address) > 0 {
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

	yptToken := &YopuNFTToken{
		AccountID:      int64(t.Account.AccountID),
		AccountAddress: t.Account.Address,
	}

	if err := yptToken.Create(); err != nil {

	}

	return nil
}

//NewMetadataURI
func (t *ERC721_YopuNFT) NewMetadataURI() (string, error) {
	return ifps.UpdateMetadataWithJSON(t)
}
