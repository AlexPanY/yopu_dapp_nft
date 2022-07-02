package token

import (
	"fmt"

	"ypt_server/config"
	"ypt_server/pkg/mysql"

	"gorm.io/gorm"
)

type YPTAccount struct {
	ID        int64 `gorm:"primaryKey"`
	Nickname  string
	Avatar    string
	Address   string
	IsDeleted int
}

func (a *YPTAccount) TableName() string {
	return `yp_account`
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

//FindByID
func (a *YPTAccount) FindByID() error {
	return a.findByMap(map[string]interface{}{
		"id": a.ID,
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
