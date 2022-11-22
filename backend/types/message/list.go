package message

type InboundEvent string
type OutboundEvent string

const (
	ConnectionSwitchEvent  OutboundEvent = "general/switch"
	EchoEvent              OutboundEvent = "general/echo"
	CardCountdown          OutboundEvent = "cd/countdown"
	CardOpen               OutboundEvent = "cd/open"
	CardState              OutboundEvent = "cd/state"
	CardDismiss            OutboundEvent = "cd/dismiss"
	LeaderboardPreview     OutboundEvent = "lb/preview"
	LeaderboardPreviewAdd  OutboundEvent = "lb/preview/add"
	LeaderboardRanking     OutboundEvent = "lb/rankings"
	LeaderboardHighlighted OutboundEvent = "lb/highlighted"
	LeaderboardPodium      OutboundEvent = "lb/podium"
	StudentTurn            OutboundEvent = "st/turn"
)
