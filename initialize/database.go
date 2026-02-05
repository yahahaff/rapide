package initialize

import (
	"errors"
	"fmt"
	"time"

	"rapide/internal/models/ssl"
	"rapide/internal/models/sys"
	"rapide/pkg/app"
	"rapide/pkg/config"
	"rapide/pkg/database"
	"rapide/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.GetString("DB_DRIVER", "sqlite") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.GetString("DB_CONNECTION_USERNAME", "root"),
			config.GetString("DB_CONNECTION_PASSWORD", "password"),
			config.GetString("DB_CONNECTION_HOST", "localhost"),
			config.GetInt("DB_CONNECTION_PORT", 3306),
			config.GetString("DB_CONNECTION_DATABASE", "rapide"),
			config.GetString("DB_CONNECTION_CHARSET", "utf8mb4"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite   window环境下启动前set CGO_ENABLED=1
		database := config.GetString("DB_CONNECTION_FILE", "./sqlite.db")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	//database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("DB_CONNECTION_MAX_OPEN_CONNECTIONS", 100))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("DB_CONNECTION_MAX_IDLE_CONNECTIONS", 10))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("DB_CONNECTION_MAX_LIFE_SECONDS", 3600)) * time.Second)

	// IsLocal本地环境 执行自动迁移表
	if app.IsLocal() {
		// First migrate tables without foreign key constraints
		err := database.DB.AutoMigrate(&sys.Role{},
			&sys.AuditLog{}, &sys.Menu{},
			&sys.UserRole{}, &sys.RoleMenu{},
			&sys.UserDept{}, &sys.Dept{},
			&sys.User{}, &ssl.SSLCert{},
		)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
