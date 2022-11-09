package message

type InboundEvent string
type OutboundEvent string

const (
	ConnectionSwitchEvent OutboundEvent = "general/switch"
	EchoEvent             OutboundEvent = "general/echo"
	CardCountdown         OutboundEvent = "cd/countdown"
	LeaderboardState      OutboundEvent = "lb/state"
)
