package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	reqHandlers()
	http.ListenAndServe(":8080", nil)
}

func reqHandlers() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("connection created")
	if err = ws.WriteMessage(1, []byte("ku pidor")); err != nil {
		log.Println(err)
	}

	read(ws)
}

func read(conn *websocket.Conn) {
	for {
		messagetype, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(messagetype, p); err != nil {
			log.Println(err)
			return
		}
	}
}
