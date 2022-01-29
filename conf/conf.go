package conf

import "github.com/BurntSushi/toml"

type Config struct {
	Mode     string `toml:"mode"`
	JWT      `toml:"jwt"`
	Server   `toml:"server"`
	Database `toml:"database"`
}

type JWT struct {
	Secret       string `toml:"secret"`
	Issuer       string `toml:"issuer"`
	ExpireTime   uint64 `toml:"expire_time_sec"`
	UserSecret   string `toml:"user_secret"`
	SellerSecret string `toml:"seller_secret"`
}

type Server struct {
	IPAddress string `toml:"ip_address"`
}

type Database struct {
	User      string `toml:"user"`
	Password  string `toml:"password"`
	IPAddress string `toml:"ip_address"`
	DBName    string `toml:"db_name"`
	MaxIdle   int    `toml:"max_idle"`
	MaxOpen   int    `toml:"max_open"`
}

const configPath string = "./conf/conf.toml"

var AppConfig Config

func Init() {
	_, err := toml.DecodeFile(configPath, &AppConfig)
	if err != nil {
		panic(err)
	}
}
