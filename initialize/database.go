package initialize

import (
	"errors"
	"fmt"
	"github.com/yahahaff/rapide/internal/models/sys"
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
	switch config.GetString("database.driver") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.GetString("database.connection.username"),
			config.GetString("database.connection.password"),
			config.GetString("database.connection.host", "localhost"),
			config.GetString("database.connection.port", 3306),
			config.GetString("database.connection.database", "rapide"),
			config.GetString("database.connection.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		database := config.GetString("database.connection.file")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	//database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.connection.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.connection.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.connection.max_life_seconds")) * time.Second)

	//gorm 自动迁移表 rbac 迁移临时用用
	err := database.DB.AutoMigrate(&sys.User{}, &sys.Menu{}, &sys.Role{}, &sys.Dept{}, &sys.OperationLog{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
