package dao

import "github.com/jinzhu/gorm"

// 定义全局变量
var (
	DB *gorm.DB
)

// 连接数据库
func InitPostgre() (err error) {
	dsn := "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=123456 sslmode=disable  TimeZone=Asia/Shanghai"
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
