package hub

import (
	"time"

	"backend/types/database"
	"backend/types/message"
)

var skip = make(chan bool)

func StartInterval(card *database.Card) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-skip:
			goto ended
		case <-ticker.C:
			card.Duration -= 1 * time.Second
			if card.Duration <= 0 {
				goto ended
			}
			Hub.CardProjectorConn.Emit(&message.OutboundMessage{
				Event: message.CardCountdown,
				Payload: map[string]any{
					"s": card.Duration.Seconds(),
					"m": card.Duration.Minutes(),
				},
			})
		}
	}

ended:
	card.Opened = true
}
