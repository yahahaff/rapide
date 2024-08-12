package initialize

import (
	"errors"
	"fmt"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/app"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.GetString("DB_DRIVER", "mysql") {
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
		// 初始化 sqlite
		database := config.GetString("DB_CONNECTION_FILE", "../database.db")
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

	// 非Release环境开启gorm 自动迁移表
	if !app.IsRelease() {
		err := database.DB.AutoMigrate(&sys.User{}, &sys.Role{},
			&sys.OperationLog{}, &sys.Menu{},
			&sys.Permission{},
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
