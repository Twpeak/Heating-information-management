package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
主要作用是数据库自动化的一些代码
 */

// ClearTable
//@description: 清理数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段，也就是设置的超时字段名，取其值跟当前时间比较，判断是否超时) string, interval(间隔) string
//@return: error

func ClearTable(db *gorm.DB,tableName string,compareField string, interval string) error {
	if db == nil { //没有连接数据库
		return errors.New("db Cannot be empty")
	}
	//将时间字符串解析为 Duration
	//Duration 为 int64类型 意为：将两个瞬间之间的经过时间表示为 int64 纳秒计数。该表示将最大可表示持续时间限制为大约 290 年。
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {	//若间隔小于0报错
		return errors.New("parse duration < 0")
	}
	//比较字段，也就是设置的超时字段名，取其值跟当前时间比较，判断是否超时。然后清除该表数据
	return db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration)).Error
}