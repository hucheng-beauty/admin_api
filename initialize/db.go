package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"admin_api/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func DB() {
	mysqlInfo := global.ServerConfig.MySQLInfo
	global.DB = connect(mysqlDsn(mysqlInfo.Endpoint, mysqlInfo.Username, mysqlInfo.Password, mysqlInfo.Database))
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名,不自动给表名加s
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // 慢 sql 阈值
				Colorful:      true,        // 禁用彩色打印
				LogLevel:      logger.Info,
			},
		),
	})
	if err != nil {
		zap.S().Panic("failed to connect database", zap.String("dsn", dsn), zap.Error(err))
	}
	return db
}

func mysqlDsn(endpoint, username, password, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, endpoint, database)
}
