/**
 * @Author: jiangbo
 * @Description:
 * @File:  answer
 * @Version: 1.0.0
 * @Date: 2021/04/24 5:57 下午
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/** 文字描述:
	我觉得不应该wrap sql.ErrNoRows错误并将其抛给上层。
	我的理由：
        接口查询数据时，有是正常的，无也应该是正常的。都应该进入调用方正常处理流程，而不是遇到无数据就进入错误处理流程
        例如demo中，如果是将sql.ErrNoRows抛给上层，则不会走打印exist这个那一行，而是进入err!=nil内部了
	demo代码如下
 */

func dao(db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRow("select name from my_user where id = ?", id).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return name, nil
}


func main() {

	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/geek")
	if err != nil {
		log.Fatal(err)
	}

	//exist, err := dao(db, 2)   ID为2存在  id为3不存在
	exist, err := dao(db, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exist:", exist)

	defer db.Close()
}
