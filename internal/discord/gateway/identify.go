package gateway

import (
	"context"
	"runtime"

	"github.com/holedaemon/hubris/internal/discord/types"
	"nhooyr.io/websocket"
)

type identifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

var defaultProperties = &identifyProperties{
	OS:      runtime.GOOS,
	Browser: "hubris",
	Device:  "hubris",
}

type identify struct {
	Token          string              `json:"token"`
	Properties     *identifyProperties `json:"properties"`
	Compress       bool                `json:"compress,omitempty"`
	LargeThreshold int                 `json:"large_threshold,omitempty"`
	// Shard          [2]int              `json:"shard,omitempty"`
	Presence *types.Presence     `json:"presence,omitempty"`
	Intents  types.GatewayIntent `json:"intents"`
}

func (c *Client) sendIdentify(ctx context.Context, ws *websocket.Conn) error {
	idf := &identify{
		Token:      c.token,
		Properties: defaultProperties,
		Intents:    types.GatewayIntentGuildsAll,
	}

	if err := write(ctx, ws, opIdentify, idf); err != nil {
		return err
	}

	return nil
}
