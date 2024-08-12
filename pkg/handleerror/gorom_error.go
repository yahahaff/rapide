package handleerror

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

func GormError(err error) string {
	if err != nil {
		// 检查错误类型
		var e *mysql.MySQLError
		switch {
		case errors.As(err, &e):
			switch e.Number {
			case 1062:
				return "记录重复"
			case 1213:
				return "Deadlock"
			default:
				// 处理其他数据库错误
				return "数据库错误"
			}
		default:
			return "未知错误"
		}
	}
	return "" // 当 err == nil 时，返回空字符串
}
