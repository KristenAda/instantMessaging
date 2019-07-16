package web

import (
	"fmt"
	"log"
	"net/http"
)
import "golang.org/x/net/websocket"

func LongLink(){
	http.Handle("/ws",websocket.Handler(Echo))

	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		if(reply=="再次回信9哈哈哈"){
			websocket.Message.Send(ws,"请求结束的消息")
			ws.Close()
		}
		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}