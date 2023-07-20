package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func SendToSocket(addr string, requestType string, dto []CallbackDto) {
	conn, err := net.Dial("unix", addr)
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	request := Request{
		Type:      requestType,
		Callbacks: dto,
	}

	// Кодирование JSON-запроса
	requestJSON, err := json.Marshal(request)
	fmt.Println(string(requestJSON))
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON-запроса:", err)
		return
	}
	//requestJSON = requestJSON[:len(requestJSON)-1]

	// Отправка JSON-запроса серверу
	_, err = conn.Write(requestJSON)
	if err != nil {
		fmt.Println("Ошибка при отправке JSON-запроса серверу:", err)
		return
	}

	// Получение ответа от сервера
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка чтения данных от сервера:", err)
		return
	}
	buffer = buffer[:n-1]

	// Распаковка JSON-ответа
	var response Response
	err = json.Unmarshal(buffer, &response)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON-ответа:", err)
	}

	// Вывод ответа
	fmt.Println("Ответ от сервера:")
	fmt.Println(response.String())
}
