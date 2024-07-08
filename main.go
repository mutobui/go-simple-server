package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Upgrader to handle WebSocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Map to store used big integers and a mutex for synchronization
var (
	usedNumbers = make(map[string]bool)
	mutex       sync.Mutex
)

// Function to generate a unique big integer
func generateUniqueBigInt() *big.Int {
	mutex.Lock()
	defer mutex.Unlock()

	var num *big.Int
	var err error
	for {
		num, err = rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))
		if err != nil {
			log.Printf("Random generation error: %v", err)
			continue
		}
		if !usedNumbers[num.String()] {
			usedNumbers[num.String()] = true
			break
		}
	}
	return num
}

// Handler for WebSocket connections
func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection attempt from:", r.RemoteAddr)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
		conn.Close()
		log.Println("Connection closed for:", r.RemoteAddr)
	}()

	log.Println("Connection established from:", r.RemoteAddr)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			}
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received message from %s: %s", r.RemoteAddr, message)

		uniqueNum := generateUniqueBigInt()
		err = conn.WriteMessage(websocket.TextMessage, []byte(uniqueNum.String()))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
		log.Printf("Sent unique number to %s: %s", r.RemoteAddr, uniqueNum.String())
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
