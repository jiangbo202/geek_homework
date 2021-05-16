/**
 * @Author: jiangbo
 * @Description:
 * @File:  register.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 11:08 下午
 */

package request

type Register struct {
	Name      string `json:"name" validate:"required"`
	Telephone string `json:"telephone" validate:"required"`
	Password  string `json:"password" validate:"required,len=6"`
}

type Login struct {
	Telephone string `json:"telephone" validate:"required"`
	Password  string `json:"password" validate:"required,len=6"`
}