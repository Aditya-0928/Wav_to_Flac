package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"wav-to-flac-converter/handlers"

	"github.com/gorilla/websocket"
)

func TestWebSocketConnection(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handlers.HandleWebSocket))
	defer server.Close()

	// Connect to the WebSocket
	url := "ws" + server.URL[4:] // Change "http" to "ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	// Prepare a sample WAV data
	wavData := []byte{
		'R', 'I', 'F', 'F', // RIFF
		0x00, 0x00, 0x00, 0x00, // file size (placeholder)
		'W', 'A', 'V', 'E', // WAVE
		'f', 'm', 't', ' ', // fmt
		0x10, 0x00, 0x00, 0x00, // subchunk1 size (16 for PCM)
		0x01, 0x00, // audio format (1 for PCM)
		0x01, 0x00, // number of channels (1 for mono)
		0x44, 0xac, 0x00, 0x00, // sample rate (44100)
		0x88, 0x58, 0x01, 0x00, // byte rate (44100 * 1 * 16/8)
		0x02, 0x00, // block align (1 * 16/8)
		0x10, 0x00, // bits per sample (16)
		'd', 'a', 't', 'a', // data
		0x00, 0x00, 0x00, 0x00, // data size (placeholder)
	}

	// Update the data size
	dataSize := len(wavData) - 44
	wavData[40] = byte(dataSize)
	wavData[41] = byte(dataSize >> 8)
	wavData[42] = byte(dataSize >> 16)
	wavData[43] = byte(dataSize >> 24)

	// Send a message
	err = conn.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read the response
	_, flacData, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	if len(flacData) == 0 {
		t.Fatal("Expected FLAC data, got empty slice")
	}
}
