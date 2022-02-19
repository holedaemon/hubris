package gateway

import (
	"context"
	"encoding/json"
	"errors"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var reconnectStatuses = map[websocket.StatusCode]bool{
	4000: true,
	4001: true,
	4002: true,
	4003: true,
	4004: false,
	4005: true,
	4007: true,
	4008: true,
	4009: true,
	4010: false,
	4011: false,
	4012: false,
	4013: false,
	4014: false,
	-1:   true,
}

func shouldReconnect(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, context.Canceled) {
		return false
	}

	if errors.Is(err, errReconnect) {
		return true
	}

	st := websocket.CloseStatus(err)
	rec, ok := reconnectStatuses[st]
	if !ok {
		return true
	}

	return rec
}

func write(ctx context.Context, ws *websocket.Conn, op opcode, v interface{}) error {
	p := &payload{
		Op: op,
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return err
	}

	p.D = raw

	return wsjson.Write(ctx, ws, p)
}
