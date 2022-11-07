package message

type InboundMessage struct {
	Event   InboundEvent `json:"event" validate:"required"`
	Payload Raw          `json:"payload"`
}

type OutboundMessage struct {
	Event   OutboundEvent `json:"event"`
	Payload any           `json:"payload"`
}

type Raw []byte

func (r *Raw) UnmarshalJSON(data []byte) error {
	*r = data
	return nil
}
