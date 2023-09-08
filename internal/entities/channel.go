package entities

type Channel struct {
	UserId      int
	ChatId      int
	ChannelLink int
}

func (c *Channel) NewChannel(UserId, ChatId, ChannelLink int) Channel {
	return Channel{UserId: UserId, ChatId: ChatId, ChannelLink: ChannelLink}
}

type ChatType struct {
	TelegramChatType bool
	StdoutChatType   bool
}

const (
	TelegramChatType = "telegram"
	StdoutChatType   = "stdout"
)
