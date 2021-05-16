/**
 * @Author: jiangbo
 * @Description:
 * @File:  database
 * @Version: 1.0.0
 * @Date: 2021/05/16 5:41 下午
 */

package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"jiang.geek/work04/model/db_model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName := viper.GetString("data.datasource.driverName")
	host := viper.GetString("data.datasource.host")
	port := viper.GetString("data.datasource.port")
	database := viper.GetString("data.datasource.database")
	username := viper.GetString("data.datasource.username")
	password := viper.GetString("data.datasource.password")
	charset := viper.GetString("data.datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接mysql失败，err:" + err.Error())
	}
	db.AutoMigrate(&db_model.User{}) // 自动创建数据表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

