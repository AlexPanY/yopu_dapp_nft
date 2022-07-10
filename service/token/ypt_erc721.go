package token

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

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
	Field string `json:"Field"`
	Value string `json:"Value"`
}

//ERC721_YopuNFT
type ERC721_YopuNFT struct {
	Account     ether.Account             `json:"-"`
	TokenID     *big.Int                  `json:"-"`
	Name        string                    `json:"Name"`
	Description string                    `json:"Description"`
	Image       string                    `json:"Image"`
	Collection  string                    `json:"Collection"`
	Properties  []*ERC721_YopuNFTProperty `json:"Properties"`
	Price       *big.Int                  `json:"-"`
	ShowPrice   int64                     `json:"Price"`
}

//NewMetadataURI
func (t *ERC721_YopuNFT) NewMetadataURI() (string, error) {
	return ifps.UpdateMetadataWithJSON(t)
}

//MetadataURI
func (t *ERC721_YopuNFT) MetadataURI() (string, error) {
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()

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

	ipfsPreix := "https://ipfs.io/"
	idx := strings.Index(uri, ipfsPreix)
	if idx < 0 {
		return errors.New("incorrect ifps URI format")
	}
	cid := uri[idx+len(ipfsPreix):]
	testuri := "http://localhost:8080/ipfs/" + cid
	resultByte, err := proxy.Get(testuri, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resultByte, &t); err != nil {
		return err
	}

	//Get Token Price
	if err := t.SellingPrice(); err != nil {
		return err
	}

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
	defer conn.Close()

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
		return err
	}

	return nil
}

func (t *ERC721_YopuNFT) SellingPrice() error {
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	contractAddress := defaultContractAddress
	if len(config.G.Contract.Address) > 0 {
		contractAddress = config.G.Contract.Address
	}

	caller, err := contracts.NewYopuNFTCaller(ether.HexToAddress(contractAddress), conn)
	if err != nil {
		return err
	}

	price, err := caller.TokenIdToPrice(&bind.CallOpts{}, t.TokenID)
	if err != nil {
		return err
	}
	t.Price = price
	t.ShowPrice = price.Int64()
	return nil
}

//SetPrice
func (t *ERC721_YopuNFT) SetPrice(price int64) error {
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

	salePrice := big.NewInt(price)
	transaction, err := ypt.SetPrice(transactOpts, t.TokenID, salePrice, ether.HexToAddress("0"))
	if err != nil {
		return err
	}

	fmt.Println(transaction)

	return err
}
