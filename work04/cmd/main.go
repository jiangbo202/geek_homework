/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/05/16 5:13 下午
 */

package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"jiang.geek/work04/configs"
	"jiang.geek/work04/utils/data"
)

func main()  {

	configs.InitConfig()
	db := data.InitDB()
	defer db.Close()

	r := configs.Routers()

	addr := viper.GetString("server.http.addr")
	if addr != "" {
		panic(r.Run(addr))
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
