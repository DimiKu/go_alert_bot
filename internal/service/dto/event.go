package dto

import "errors"

var EmptyEventKey = errors.New("key of event can't be empty")

type EventDto struct {
	Key         string `json:"key"`
	UserId      int    `json:"user_id"`
	ChannelLink ChannelLinkDto
}

func (e *EventDto) Validate() error {
	if e.Key == "" {
		return EmptyEventKey
	}

	return nil
}
