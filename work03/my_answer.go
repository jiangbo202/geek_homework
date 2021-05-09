/**
 * @Author: jiangbo
 * @Description:
 * @File:  my_answer
 * @Version: 1.0.0
 * @Date: 2021/05/09 4:24 下午
 */

package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

// 启动了三个端口的server，如果有请求到127.0.0.1:8003/index3则会导致这三个server全部关闭

func main() {

	group := new(errgroup.Group)
	for index := 0; index < 3; index++ {
		indexTemp := index
		group.Go(func() error {
			fmt.Printf("indexTemp=%d \n", indexTemp)

			if indexTemp == 0 {
				http.HandleFunc("/index1", func(writer http.ResponseWriter, request *http.Request) {
					fmt.Fprintf(writer, "hello world %d~", indexTemp)
				})
				http.ListenAndServe("127.0.0.1:8001", nil)
			} else if indexTemp == 1 {
				http.HandleFunc("/index2", func(writer http.ResponseWriter, request *http.Request) {
					fmt.Fprintf(writer, "hello world %d~", indexTemp)
				})
				http.ListenAndServe("127.0.0.1:8002", nil)
			}else if indexTemp == 2 {
				http.HandleFunc("/index3", func(writer http.ResponseWriter, request *http.Request) {

					fmt.Fprintf(writer, "hello world %d~ 请求我，我就停所有服务了", indexTemp)
					log.Fatal("触发停所有服务(来自8003)。。。")
				})
				http.ListenAndServe("127.0.0.1:8003", nil)
			}
			return nil
		})
	}

	// 捕获err
	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}


}
