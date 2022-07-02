package token

type YopuNFTToken struct {
	ID             int64 `gorm:"primaryKey"`
	AccountID      int64
	AccountAddress string
	AccountToken   string
	is_deleted     int
}

//TableName
func (a *YopuNFTToken) TableName() string {
	return `yp_account_token`
}

//Create
func (t *YopuNFTToken) Create() error {
	return mysql.DB[config.DB_YPT].Create(t).Error
}
