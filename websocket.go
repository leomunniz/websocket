package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/leomunniz/websocket/core/mongo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Println("Received:", string(msg))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", handler)
	mongo.Connect_mongo_db()

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

}
