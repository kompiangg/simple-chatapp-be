package impl

import (
	"github.com/gorilla/websocket"
	pkgWebsocket "github.com/kompiangg/chatapp/pkg/websocket"
)

type WebSocketImpl struct {
	pool *pkgWebsocket.Pool
}

func (w WebSocketImpl) CreateClient(conn *websocket.Conn) {
	client := &pkgWebsocket.Client{
		Conn: conn,
		Pool: w.pool,
	}

	w.pool.Register <- client
	client.Read()
}

func ProvideWebSocketService(pool *pkgWebsocket.Pool) *WebSocketImpl {
	return &WebSocketImpl{
		pool: pool,
	}
} 