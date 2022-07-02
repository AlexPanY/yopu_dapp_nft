package config

type ContractConf struct {
	Address   string `json:"address"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	NetworkID int    `json:"network_id"`
	Gas       int    `json:"gas"`
	GasPrice  int    `json:"gas_price"`
	From      string `json:"from"`
}
