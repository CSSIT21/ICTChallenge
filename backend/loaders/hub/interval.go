package hub

import (
	"math"
	"time"

	"backend/types/database"
	"backend/types/message"
)

var Skip = make(chan bool)
var Pause = make(chan bool)

func StartInterval(card *database.Card) {
	ticker := time.NewTicker(1 * time.Second)

	Hub.CardProjectorConn.Emit(&message.OutboundMessage{
		Event: message.CardCountdown,
		Payload: map[string]any{
			"s": int(math.Floor(card.Duration.Seconds())) % 60,
			"m": math.Floor(card.Duration.Minutes()),
		},
	})

	for {
		select {
		case <-Skip:
			goto ended
		case <-Pause:
			<-Pause
		case <-ticker.C:
			card.Duration -= 1 * time.Second
			if card.Duration <= 0 {
				goto ended
			}
			Hub.CardProjectorConn.Emit(&message.OutboundMessage{
				Event: message.CardCountdown,
				Payload: map[string]any{
					"s": int(math.Floor(card.Duration.Seconds())) % 60,
					"m": math.Floor(card.Duration.Minutes()),
				},
			})
		}
	}

ended:
	Hub.CardProjectorConn.Emit(&message.OutboundMessage{
		Event: message.CardCountdown,
		Payload: map[string]any{
			"s": 0,
			"m": 0,
		},
	})
	card.Opened = true
}
