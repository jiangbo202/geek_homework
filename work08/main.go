/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/06/14 6:41 下午
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"unsafe"
)

func main()  {

	buf := [5000]byte{}
	size := 0

	size, _ = push(buf[:], size, 'a')
	fmt.Printf("a的字节数是%d", unsafe.Sizeof(buf))

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// 循环一w次
	for i := 0; i < 10000; i++ {
		_, err = c.Do("SET", "mykey"+string(i), buf)
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	}
	fmt.Println("over")
}

func push(buf []byte, size int, b byte) (int, error) {
	max := len(buf)

	if max < 1 {
		return size, fmt.Errorf("buffer underflow: max=%d char=%d", max, b)
	}

	if size >= max {
		return size, fmt.Errorf("buffer overflow: size=%d max=%d char=%d", size, max, b)
	}

	buf[size] = b

	return size + 1, nil
}