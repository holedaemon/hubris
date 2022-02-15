package gateway

import (
	"encoding/json"

	"github.com/holedaemon/hubris/internal/discord/types/null"
)

type payload struct {
	Op opcode          `json:"op"`
	D  json.RawMessage `json:"d"`
	S  null.Int        `json:"s"`
	T  null.String     `json:"t"`
}

type hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}
