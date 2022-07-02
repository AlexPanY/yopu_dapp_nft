package ether

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	AccountID        uint64
	Address          string
	EthAddress       common.Address
	PrivateKey       string
	ECDSA_PublicKey  *ecdsa.PublicKey  `json:"-"`
	ECDSA_PrivateKey *ecdsa.PrivateKey `json:"-"`
}

//NewAccount
func NewAccount(accountID uint64, address, key string) Account {
	a := Account{
		AccountID:  accountID,
		Address:    address,
		EthAddress: common.HexToAddress(address),
		PrivateKey: key,
	}
	a.ECDSAKeyWithPrivateKey()
	return a
}

//HexToAddress
func HexToAddress(address string) common.Address {
	return common.HexToAddress(address)
}

//GetAccountByID
func GetAccountByID(accountID uint64) Account {
	a := Account{
		AccountID: accountID,
		Address:   "0x71c7656ec7ab88b098defb751b7401b5f6d8976f",
	}
	a.ECDSAKeyWithPrivateKey()
	a.EthAddress = common.HexToAddress(a.Address)
	return a
}

//ParseECDSAKeyWithPrivateKey
func (a *Account) ECDSAKeyWithPrivateKey() error {
	//1.解析私钥
	privateKey, err := crypto.HexToECDSA(a.PrivateKey)
	if err != nil {
		return err
	}
	a.ECDSA_PrivateKey = privateKey

	//2.获取发送账户的公共地址
	//NOTE: Public Address 可以通过私钥派生出来
	publicKey := a.ECDSA_PrivateKey.Public()
	publicKeyECDSA, isOk := publicKey.(*ecdsa.PublicKey)
	if !isOk {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	a.ECDSA_PublicKey = publicKeyECDSA
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(fromAddress)
	return nil
}

//TransferETHByAddress 通过ETH进行转账交易
func (a *Account) TransferETHByAddress(toAdderss string, amount int64) error {
	err := Transfer(a.PrivateKey, toAdderss, amount)
	if err != nil {
		return err
	}
	return nil
}

//GetBalanceByAddress
func (a *Account) GetBalanceByAddress() (*big.Int, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), a.EthAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

//GetPendingBalanceByAddress 获取账户待处理余额
func (a *Account) GetPendingBalanceByAddress() (*big.Int, error) {
	client, err := NewEtherClientConn()
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), a.EthAddress, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
