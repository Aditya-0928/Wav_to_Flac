package handlers

import (
	"log"
	"net/http"
	"wav-to-flac-converter/audio"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, wavData, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		flacData, err := audio.ConvertWAVToFLAC(wavData)
		if err != nil {
			log.Println("Conversion error:", err)
			continue
		}

		if err = conn.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
