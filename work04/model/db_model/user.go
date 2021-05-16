/**
 * @Author: jiangbo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/01/02 4:16 下午
 */

package db_model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
