package configs

import "github.com/DropKit/DropKit-Adapter/pkg/logging"

type Configurations struct {
	Quorum     QuorumConfigs
	YugabyteDB YugabyteDBConfigs
	Dropkit    DropkitConfigs
	Price      PriceConfigs
	Adapter    AdapterConfigs
}

type QuorumConfigs struct {
	Endpoint string
}

type YugabyteDBConfigs struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type DropkitConfigs struct {
	Account  string
	Contract string
	Host     string
	Port     string
}

type PriceConfigs struct {
	Create int
	Insert int
	Delete int
	Select int
	Update int
}

type AdapterConfigs struct {
	LoggerLevel logging.LoggerLevel
}

var Configs Configurations
