package cli

import (
	"fmt"
	"net"
)

func SendToSocket(addr string, requestType string, message []byte) ([]byte, error) {
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

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка чтения данных от сервера:", err)
		return nil, err
	}
	buffer = buffer[:n]
	return buffer, nil
}
