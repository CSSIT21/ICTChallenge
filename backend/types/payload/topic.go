package payload

type OpenCard struct {
	TopicId uint64 `json:"topic_id" validate:"required"`
	CardId  uint64 `json:"card_id" validate:"required"`
}
