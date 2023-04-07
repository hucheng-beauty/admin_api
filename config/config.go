package config

type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	Port      string      `mapstructure:"port" json:"port"`
	MySQLInfo MySQLConfig `mapstructure:"mysql_info" json:"mysql_info"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
}

type MySQLConfig struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
