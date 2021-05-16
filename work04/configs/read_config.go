/**
 * @Author: jiangbo
 * @Description:
 * @File:  read_config
 * @Version: 1.0.0
 * @Date: 2021/05/16 5:34 下午
 */

package configs

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
