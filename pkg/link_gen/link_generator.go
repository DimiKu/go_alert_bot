package link_gen

import (
	"math/rand"
)

type ChannelLink int64

func LinkGenerate() ChannelLink {
	r := rand.New(rand.NewSource(999))
	return ChannelLink(r.Int63())

}
