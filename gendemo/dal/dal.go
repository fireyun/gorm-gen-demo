package dal

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
		// DB = ConnectDB().Debug()
		// _ = DB.AutoMigrate(&model.User{}, &model.Passport{}, &model.TemplateSets{})
		// _ = DB.AutoMigrate(&model.TemplateSets{})
	})
}

func ConnectDB() (conn *gorm.DB) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bk_bscp_admin?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
