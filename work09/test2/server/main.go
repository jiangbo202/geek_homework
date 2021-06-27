/**
 * @Author: jiangbo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/06/27 9:13 下午
 */

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {

	netListen, err := net.Listen("tcp", ":9988")

	CheckError(err)



	defer netListen.Close()



	Log("Waiting for clients")

	for {

		conn, err := netListen.Accept()

		if err != nil {

			continue

		}



		Log(conn.RemoteAddr().String(), " tcp connect success")

		go handleConnection(conn)

	}

}



func handleConnection(conn net.Conn) {

	//声明一个临时缓冲区，用来存储被截断的数据

	tmpBuffer := make([]byte, 0)



	//声明一个管道用于接收解包的数据

	readerChannel := make(chan []byte, 16)

	go reader(readerChannel)



	buffer := make([]byte, 1024)

	for {

		n, err := conn.Read(buffer)

		if err != nil {

			Log(conn.RemoteAddr().String(), " connection error: ", err)

			return

		}



		tmpBuffer = Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)

	}

}



func reader(readerChannel chan []byte) {

	for {

		select {

		case data := <-readerChannel:

			Log(string(data))

		}

	}

}



func Log(v ...interface{}) {

	fmt.Println(v...)

}



func CheckError(err error) {

	if err != nil {

		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

		os.Exit(1)

	}

}

const (
	ConstHeader         = "www.01happy.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			readerChannel <- data
			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

//整形转换成字节

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)

}
