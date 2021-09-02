package main

import (
	"log"
	"net/http"
	"video-chat-app/controllers"
)

func main() {
	http.Handle("/create", controllers.CreateRoomRequestHandler)
	http.Handle("/join", controllers.JoinRoomRequestHandler)

	log.Println("Starting on port 7000")
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
