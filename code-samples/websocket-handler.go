// Early WebSocket Implementation - WeWatch Platform
// This shows concurrent connection handling with Go channels and goroutines

package handlers

import (
    "encoding/json"
    "sync"
    "github.com/gorilla/websocket"
)

// WebSocket message structure
type WebSocketMessage struct {
    Type        string      `json:"type"`
    Command     string      `json:"command,omitempty"`
    UserID      uint        `json:"user_id,omitempty"`
    RoomID      uint        `json:"room_id,omitempty"`
    Data        interface{} `json:"data,omitempty"`
}

// Client represents a single WebSocket connection
type Client struct {
    hub      *Hub
    conn     *websocket.Conn
    send     chan []byte      // Buffered channel for outbound messages
    roomID   uint              // Room subscription
    userID   uint              // Authenticated user
}

// Hub manages all active connections and message routing
type Hub struct {
    // Map of rooms to their connected clients
    rooms map[uint]map[*Client]bool
    
    // Channels for concurrent message handling
    broadcast    chan []byte          // Broadcast to all
    register     chan *Client         // New client connections
    unregister   chan *Client         // Client disconnections
    
    // Mutex for thread-safe map access
    mutex sync.RWMutex
}

// Run starts the hub's main event loop (goroutine)
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            // Register new client to their room
            h.mutex.Lock()
            if h.rooms[client.roomID] == nil {
                h.rooms[client.roomID] = make(map[*Client]bool)
            }
            h.rooms[client.roomID][client] = true
            h.mutex.Unlock()
            
        case client := <-h.unregister:
            // Remove client from room
            h.mutex.Lock()
            if clients, ok := h.rooms[client.roomID]; ok {
                if _, exists := clients[client]; exists {
                    delete(clients, client)
                    close(client.send)
                }
            }
            h.mutex.Unlock()
            
        case message := <-h.broadcast:
            // Broadcast message to all clients in all rooms
            h.mutex.RLock()
            for _, clients := range h.rooms {
                for client := range clients {
                    select {
                    case client.send <- message:
                    default:
                        // Client can't receive, remove it
                        close(client.send)
                        delete(clients, client)
                    }
                }
            }
            h.mutex.RUnlock()
        }
    }
}

// BroadcastToRoom sends message to specific room
func (h *Hub) BroadcastToRoom(roomID uint, message []byte) {
    h.mutex.RLock()
    defer h.mutex.RUnlock()
    
    if clients, ok := h.rooms[roomID]; ok {
        for client := range clients {
            select {
            case client.send <- message:
            default:
                // Channel full, skip this client
            }
        }
    }
}

// Key Learning: Go's channels + goroutines = elegant real-time systems
// - Channels act as thread-safe queues
// - select statement handles multiple channel operations
// - RWMutex allows concurrent reads, exclusive writes
