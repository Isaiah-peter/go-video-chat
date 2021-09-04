package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"video-chat-app/models"
)

var AllRooms models.RoomMap

type resp struct {
	RoomId string `json:"room_id"`
}

// create a room and return a
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId := AllRooms.CreateRoom()
	json.NewEncoder(w).Encode(resp{RoomId: roomId})
}

//join room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello word")
}
