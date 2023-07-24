package cli

import (
	"fmt"
	"io/ioutil"
	"net"
)

func SendToSocket(addr string, message []byte) ([]byte, error) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return nil, err
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
	}(conn)

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Ошибка при отправке JSON-запроса серверу:", err)
		return nil, err
	}
	var buff []byte
	buff, err = ioutil.ReadAll(conn)
	fmt.Println(string(buff))
	return buff, err
}
