package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader  = websocket.Upgrader{}

type Message struct {
	Email 		string `json:"email"`
	UserName 	string `json:username`
	Message 	string `json:message`
}

func main() {

	//create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	//configure websocket route
	http.HandleFunc("/ws", handleConnections)

	//start listening for incoming chat message
	go handleMessages()

	//start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("listenAndServe", err)
	}
}

func handleConnections(writer http.ResponseWriter,request *http.Request) {
	ws,err := upgrader.Upgrade(writer,request, nil)
	if err != nil{
		log.Fatal(err)
	}

	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		err:= ws.ReadJSON(&msg)
		if err != nil{
			log.Printf("error:%v",err)
			delete(clients,ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages(){
	for{
		msg := <-broadcast
		for client := range clients{
			err := client.WriteJSON(msg)
			if err != nil{
				log.Printf("error:%v",err)
				client.Close()
				delete(clients,client)
			}
		}
	}
}
