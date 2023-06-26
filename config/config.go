package config

type ServerConfig struct {
	Name        string      `mapstructure:"name" json:"name"`
	Host        string      `mapstructure:"host" json:"host"`
	Port        string      `mapstructure:"port" json:"port"`
	MySQLInfo   MySQLConfig `mapstructure:"mysql_info" json:"mysql_info"`
	RedisInfo   RedisConfig `mapstructure:"redis_info" json:"redis_info"`
	JWTInfo     JWTConfig   `mapstructure:"jwt" json:"jwt"`
	EmailConfig EmailConfig `mapstructure:"email_info" json:"email_info"`
}

type MySQLConfig struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type RedisConfig struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
	Password string `mapstructure:"username" json:"username"`
	Database int    `mapstructure:"password" json:"password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type EmailConfig struct {
	EmailID  string `mapstructure:"emailID" json:"emailID"`
	Password string `mapstructure:"password" json:"password"`
	SmtpHost string `mapstructure:"smtp_host" json:"smtp_host"`
	SmtpPort string `mapstructure:"smtp_port" json:"smtp_port"`
}
