package token

import (
	"context"
	"fmt"
	"math/big"

	"ypt_server/config"
	"ypt_server/contracts"
	"ypt_server/pkg/ether"
	"ypt_server/pkg/mysql"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gorm.io/gorm"
)

type YPTAccount struct {
	ID         int64 `gorm:"primaryKey"`
	Nickname   string
	Avatar     string
	Address    string
	Privatekey string `gorm:"-"`
	IsDeleted  int
}

func (a *YPTAccount) TableName() string {
	return `yp_account`
}

//Buy
func (a *YPTAccount) SafeTransfer(to string, iTokenID int64) error {

	tokenID := big.NewInt(iTokenID)
	t := ERC721_YopuNFT{
		Account: ether.NewAccount(1, a.Address, a.Privatekey),
		TokenID: tokenID,
	}

	if err := t.Get(); err != nil {
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

	ts, err := ypt.SafeTransferFrom(transactOpts, t.Account.EthAddress, ether.HexToAddress(to), tokenID)
	if err != nil {
		return err
	}

	fmt.Println(ts)

	return nil
}

//BalanceOf
func (a *YPTAccount) BalanceOf() error {
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return err
	}

	contractAddress := defaultContractAddress
	if len(config.G.Contract.Address) > 0 {
		contractAddress = config.G.Contract.Address
	}

	caller, err := contracts.NewYopuNFTCaller(ether.HexToAddress(contractAddress), conn)
	if err != nil {
		return err
	}

	callopts := &bind.CallOpts{}

	count, err := caller.BalanceOf(callopts, ether.HexToAddress(a.Address))
	if err != nil {
		return err
	}

	var i int64 = 1
	for {
		tokenID := big.NewInt(i)

		if count.Cmp(tokenID) < 0 {
			break
		}

		tokenURI, err := caller.TokenURI(callopts, tokenID)
		if err != nil {
			break
		}
		fmt.Println(fmt.Sprintf("------TokenID: %v, URI: %v-----", tokenID, tokenURI))

		i++
	}

	fmt.Println(fmt.Sprintf("------BalanceOf: %v-----", count))
	return nil
}

//Create
func (a *YPTAccount) Create() error {
	return mysql.DB[config.DB_YPT].Create(a).Error
}

//FindAccountByID
func FindAccountByID(accountID int64) (a YPTAccount, err error) {
	a = YPTAccount{
		ID: accountID,
	}
	err = (&a).FindByID()
	return
}

//FindAccountByAddress
func FindAccountByAddress(address string) (a YPTAccount, err error) {
	a = YPTAccount{
		Address: address,
	}
	err = (&a).FindByAddress()
	return
}

//FindByID
func (a *YPTAccount) FindByID() error {
	return a.findByMap(map[string]interface{}{
		"id": a.ID,
	})
}

//FindByID
func (a *YPTAccount) FindByAddress() error {
	return a.findByMap(map[string]interface{}{
		"address": a.Address,
	})
}

func (a *YPTAccount) findByMap(wheremaps map[string]interface{}) error {
	if err := mysql.DB[config.DB_YPT].Table(a.TableName()).Where(wheremaps).Find(a).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}

//FindAccountAddressByID
func FindAccountAddressByID(accountID int64) (address string) {
	sql := fmt.Sprintf("SELECT `address` FROM `yp_account` WHERE id = %d;",
		accountID,
	)
	mysql.DB[config.DB_YPT].Raw(sql).Scan(&address)
	return
}
