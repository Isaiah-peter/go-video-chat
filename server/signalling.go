package server

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

type resp struct {
	RoomID string `json:"room_id"`
}

type broadcastMsg struct {
	 Message map[string]interface{}
	 RoomID  string
	 Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster()  {
   for {
	   msg := <- broadcast

	   for _, client := range AllRooms.Map[msg.RoomID] {
		   if client.Conn != msg.Client {
			   err := client.Conn.WriteJSON(msg.Message)

			   if err != nil {
				   log.Fatal(err)
				   client.Conn.Close()
			   }
		   }
	   }
   }
}

// CreateRoomRequestHandler create a room and return a
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	roomId := AllRooms.CreateRoom()
	log.Println(AllRooms.Map)
    json.NewEncoder(w).Encode(resp{RoomID: roomId })
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// JoinRoomRequestHandler join room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId, ok := r.URL.Query()["roomID"]
	if !ok {
		log.Println("roomID messing in URL parameter")
		return
	}

	ws, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal("Websocket upgrade Err", err)
	}

	AllRooms.InsertIntoRoom(roomId[0], false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(msg.Message)
		if err != nil {
			log.Fatal("Read err: ", err)
		}
		msg.Client = ws
		msg.RoomID = roomId[0]

		broadcast <- msg
	}

}
