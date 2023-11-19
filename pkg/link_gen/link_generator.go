package link_gen

import (
	"math/rand"
	"time"
)

type ChannelLink int64

func LinkGenerate() ChannelLink {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return ChannelLink(r.Int63())
}
