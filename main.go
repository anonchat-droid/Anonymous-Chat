package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // Upgrader for WebSocket

var clients = make(map[*websocket.Conn]bool) // Connected clients

func handleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Connection error:", err)
        return
    }
    clients[conn] = true
    defer func() {
        delete(clients, conn)
        conn.Close()
    }()

    for {
        var message string
        err := conn.ReadJSON(&message)
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }
        broadcast(message)
    }
}

func broadcast(message string) {
    for client := range clients {
        err := client.WriteJSON(message)
        if err != nil {
            fmt.Println("Error sending message:", err)
            client.Close()
            delete(clients, client)
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnection)
    fmt.Println("Server is running on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
