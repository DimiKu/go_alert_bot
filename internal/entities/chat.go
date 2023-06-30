package entities

type Chat struct {
	UserId int
	ChatId int
}

func (c *Chat) NewChat(UserId, ChatId int) Chat {
	return Chat{UserId: UserId, ChatId: ChatId}
}
