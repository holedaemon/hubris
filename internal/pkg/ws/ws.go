package ws

import (
	"bytes"
	"compress/zlib"
	"context"
	"fmt"
	"io"
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
func (ws *Conn) Read(ctx context.Context) (io.Reader, error) {
	ws.mu.RLock()
	mt, rdr, err := ws.conn.Reader(ctx)
	ws.mu.RUnlock()

	if err != nil {
		return nil, fmt.Errorf("getting reader from conn: %w", err)
	}

	if mt == websocket.MessageText {
		return rdr, nil
	}

	zr, err := zlib.NewReader(rdr)
	if err != nil {
		return nil, fmt.Errorf("decompressing reader: %w", err)
	}

	var out bytes.Buffer

	_, err = io.Copy(&out, zr)
	if err != nil {
		return nil, fmt.Errorf("copying reader: %w", err)
	}
	zr.Close()
	return &out, nil
}

// Write writes to a connection.
func (ws *Conn) Write(ctx context.Context, v interface{}) error {
	return wsjson.Write(ctx, ws.conn, v)
}
