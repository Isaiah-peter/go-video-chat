package main

import (
	"log"
	"net/http"
	"video-chat-app/server"
)

func main() {
	server.AllRooms.Init()
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting on port 7000")
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
