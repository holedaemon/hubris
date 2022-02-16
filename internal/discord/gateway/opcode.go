package gateway

type opcode int

const (
	opDispatch opcode = iota
	opHeartbeat
	opIdentify
	opPresenceUpdate   //nolint:deadcode
	opVoiceStateUpdate //nolint:deadcode
	_
	opResume
	opReconnect
	opRequestGuildMembers //nolint:deadcode
	opInvalidSession
	opHello
	opHeartbeatAck
)
