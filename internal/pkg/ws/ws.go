package ws

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Conn wraps a websocket connection for thread-safety.
type Conn struct {
	mu   sync.RWMutex
	conn *websocket.Conn
}

// Dial opens a connection to u with the given headers.
func Dial(ctx context.Context, u string, h http.Header) (*Conn, error) {
	ws, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		HTTPHeader: h,
	})
	if err != nil {
		return nil, err
	}

	return &Conn{conn: ws}, nil
}

// Close closes a connection.
func (ws *Conn) Close(st websocket.StatusCode, m string) error {
	return ws.conn.Close(st, m)
}

// Read reads from a connection.
func (ws *Conn) Read(ctx context.Context, v interface{}) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	typ, raw, err := ws.conn.Read(ctx)
	if err != nil {
		return err
	}

	in := bytes.NewBuffer(raw)

	if typ == websocket.MessageBinary {
		zr, err := zlib.NewReader(in)
		if err != nil {
			return err
		}

		if err := json.NewDecoder(zr).Decode(&v); err != nil {
			return err
		}

		zr.Close()

		return nil
	}

	if err := json.NewDecoder(in).Decode(&v); err != nil {
		return err
	}

	return err
}

// Write writes to a connection.
func (ws *Conn) Write(ctx context.Context, v interface{}) error {
	return wsjson.Write(ctx, ws.conn, v)
}
