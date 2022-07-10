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
func (a *YPTAccount) Buy(iTokenID int64) error {
	tokenID := big.NewInt(iTokenID)

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

	buyerAccount := ether.NewAccount(1, a.Address, a.Privatekey)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(buyerAccount.ECDSA_PrivateKey, chainID)
	if err != nil {
		return err
	}

	t := ERC721_YopuNFT{
		TokenID: tokenID,
	}
	if err := t.SellingPrice(); err != nil {
		return err
	}
	transactOpts.Value = t.Price

	ts, err := ypt.Buy(transactOpts, tokenID)
	if err != nil {
		return err
	}
	fmt.Println(ts)
	return nil
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
	if err := t.SellingPrice(); err != nil {
		return err
	}
	transactOpts.Value = t.Price
	ts, err := ypt.SafeTransferFrom(transactOpts, t.Account.EthAddress, ether.HexToAddress(to), tokenID)
	if err != nil {
		return err
	}

	fmt.Println(ts)

	return nil
}

//BalanceOf
func (a *YPTAccount) BalanceOf() (*big.Int, error) {
	conn, err := ether.NewEtherClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	contractAddress := defaultContractAddress
	if len(config.G.Contract.Address) > 0 {
		contractAddress = config.G.Contract.Address
	}

	caller, err := contracts.NewYopuNFTCaller(ether.HexToAddress(contractAddress), conn)
	if err != nil {
		return nil, err
	}

	callopts := &bind.CallOpts{}

	count, err := caller.BalanceOf(callopts, ether.HexToAddress(a.Address))
	if err != nil {
		return nil, err
	}
	return count, err

}

//Tokens All tokens under the account
func (a *YPTAccount) Tokens() ([]*ERC721_YopuNFT, error) {
	count, err := a.BalanceOf()
	if err != nil {
		return nil, err
	}

	var i int64 = 0
	erc721Tokens := make([]*ERC721_YopuNFT, 0, 2)
	for {
		tokenID := big.NewInt(i)
		if tokenID.Cmp(count) >= 0 {
			break
		}

		t := &ERC721_YopuNFT{
			TokenID: tokenID,
		}
		if err := t.Get(); err != nil {
			continue
		}
		erc721Tokens = append(erc721Tokens, t)
		i++
	}

	return erc721Tokens, nil
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
