package message

type InboundEvent string
type OutboundEvent string

const (
	ConnectionSwitchEvent OutboundEvent = "general/switch"
	EchoEvent             OutboundEvent = "general/echo"
	CardCountdown         OutboundEvent = "cd/countdown"
	CardOpen              OutboundEvent = "cd/open"
	CardState             OutboundEvent = "cd/state"
	LeaderboardState      OutboundEvent = "lb/state"
	LeaderboardPodium     OutboundEvent = "lb/podium"
	StudentTurn           OutboundEvent = "st/turn"
)
