package gateway

type opcode int

//nolint
const (
	opDispatch opcode = iota
	opHeartbeat
	opIdentify
	opPresenceUpdate
	opVoiceStateUpdate
	_
	opResume
	opReconnect
	opRequestGuildMembers
	opInvalidSession
	opHello
	opHeartbeatAck
)
