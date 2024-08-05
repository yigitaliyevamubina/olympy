package websocket

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	clients  = make(map[string]*websocket.Conn)
)

// JWT secret key (replace with your actual secret key)
var jwtSecret = []byte("thisisaverysecretsecretkeywrittenbynodirbek")

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Extract and verify the JWT token from the request
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		http.Error(w, "token is required", http.StatusUnauthorized)
		return
	}

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade connection: %v", err)
		return
	}

	// Store the authenticated client connection
	clients[conn.RemoteAddr().String()] = conn

	defer func() {
		conn.Close()
		delete(clients, conn.RemoteAddr().String())
	}()

	// Listen for messages from the client
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
	}
}

func Broadcast(event interface{}) {
	for _, client := range clients {
		err := client.WriteJSON(event)
		if err != nil {
			log.Printf("failed to send message to client: %v", err)
			client.Close()
			delete(clients, client.RemoteAddr().String())
		}
	}
}
