package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	sockerUrl := "ws://localhost:8080" + "/ws"
	// 客户端结束信号
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	// 关闭websocket信号
	done := make(chan struct{})
	// websocket连接
	conn, _, err := websocket.DefaultDialer.Dial(sockerUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		defer close(done)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("[websocket client] received:", string(msg))
		}
	}()

	go func() {
		select {
		case sig := <-interrupt:
			log.Println("received sigint interrupt signal. Closing all connections", sig)
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "end..."))
			if err != nil {
				log.Println(err)
			}
			select {
			case <-done:
				log.Println("close websocket conn")
			case <-time.After(time.Second * 3):
				log.Println("timeout in closing websocket conn")
			}
			conn.Close()
			os.Exit(1)
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("this is a test msg."))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
