package token

import (
	"fmt"
	"ypt_server/config"
	"ypt_server/pkg/mysql"
)

type YopuNFTToken struct {
	ID             int64 `gorm:"primaryKey"`
	AccountID      int64
	AccountAddress string
	AccountToken   string
	IsDeleted      int
}

//TableName
func (a *YopuNFTToken) TableName() string {
	return `yp_account_token`
}

//Create
func (t *YopuNFTToken) Create() error {
	return mysql.DB[config.DB_YPT].Create(t).Error
}

//FindTokensByAccountID
func FindTokensByAccountID(accountID int64) (tokens []*YopuNFTToken) {
	sql := fmt.Sprintf("SELECT * FROM `yp_account_token` WHERE account_id = %d;",
		accountID,
	)

	mysql.DB[config.DB_YPT].Raw(sql).Scan(&tokens)
	if len(tokens) <= 0 {
		tokens = make([]*YopuNFTToken, 0)
	}

	return
}
