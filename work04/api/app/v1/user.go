/**
 * @Author: jiangbo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/05/16 8:05 下午
 */

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"jiang.geek/work04/model/db_model"
	"jiang.geek/work04/model/request"
	"jiang.geek/work04/model/response"
	"jiang.geek/work04/service"
	"net/http"
)

func Register(c *gin.Context) {
	var register request.Register
	c.ShouldBindJSON(&register)

	// 校验register
	fmt.Println(register.Name, register.Password, register.Telephone)
	validate := validator.New()
	errs := validate.Struct(register)
	if errs != nil {
		fmt.Println(errs.Error())
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, errs.Error())
		return
	}

	// 手机号位数验证
	if len(register.Telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	// 密码长度校验
	if len(register.Password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	token, err := service.Register(&register)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token生成错误: "+ err.Error())
		return
	}
	//
	response.Success(c, gin.H{"token": token}, "注册成功")
}


func Login(c *gin.Context) {
	// 获取参数
	var login request.Login
	c.ShouldBindJSON(&login)

	fmt.Println(login.Password, login.Telephone)
	validate := validator.New()
	errs := validate.Struct(login)
	if errs != nil {
		fmt.Println(errs.Error())
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, errs.Error())
		return
	}

	// 手机号长度检验
	if len(login.Telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 密码长度检验
	if len(login.Password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	token, err :=service.Login(&login)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, err.Error())
		return
	}
	// 返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	fmt.Println("user:", user)
	response.Success(c, gin.H{"user": response.ToUserResponse(user.(db_model.User))}, "登陆成功")
}