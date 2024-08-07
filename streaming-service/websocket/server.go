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

var jwtSecret = []byte("nodirbek")

func HandleConnections(w http.ResponseWriter, r *http.Request) {
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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade connection: %v", err)
		return
	}

	clients[conn.RemoteAddr().String()] = conn

	defer func() {
		conn.Close()
		delete(clients, conn.RemoteAddr().String())
	}()

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
