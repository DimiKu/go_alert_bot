package dto

type EventDto struct {
	Key         string `json:"key"`
	UserId      int    `json:"user_id"`
	ChannelLink ChannelLinkDto
}
