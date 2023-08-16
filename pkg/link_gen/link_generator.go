package link_gen

import (
	"math/rand"
	"time"
)

type ChannelLink int64

func LinkGenerate() ChannelLink {
	rand.Seed(time.Now().UnixNano())
	return ChannelLink(rand.Int63())

}
