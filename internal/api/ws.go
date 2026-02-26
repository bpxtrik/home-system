package api

import (
	"home-system/internal"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true }, // adjust in prod
}

var clients = struct {
    sync.RWMutex
    conns map[*websocket.Conn]bool
}{conns: make(map[*websocket.Conn]bool)}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }

    clients.Lock()
    clients.conns[conn] = true
    clients.Unlock()

    // keep connection open
    for {
        if _, _, err := conn.NextReader(); err != nil {
            clients.Lock()
            delete(clients.conns, conn)
            clients.Unlock()
            conn.Close()
            break
        }
    }
}

func broadcastMotion(m internal.Motion) {
    clients.RLock()
    defer clients.RUnlock()

    for conn := range clients.conns {
        conn.WriteJSON(m)
    }
}
