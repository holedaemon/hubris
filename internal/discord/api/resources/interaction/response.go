package interaction

import "github.com/holedaemon/hubris/internal/discord/types"

func NewMessageResponse(m *types.Message) *types.InteractionResponse {
	return &types.InteractionResponse{
		Type: types.InteractionCallbackTypeChannelMessageWithSource,
		Data: m,
	}
}

func NewDeferredMessageResponse(m *types.Message) *types.InteractionResponse {
	return &types.InteractionResponse{
		Type: types.InteractionCallbackTypeDeferredChannelMessageWithSource,
		Data: m,
	}
}
