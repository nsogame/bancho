package packets

import "github.com/Mempler/osubinary"

type ChannelJoinSuccessPacket struct {
	name string
}

func (ChannelJoinSuccessPacket) GetID() uint16 {
	return BanchoLoginPermissions
}

func (p ChannelJoinSuccessPacket) GetData() []byte {
	return osubinary.BString(p.name)
}

func ChannelJoinSuccess(name string) ChannelJoinSuccessPacket {
	return ChannelJoinSuccessPacket{name}
}
