package packets

import "github.com/Mempler/osubinary"

type LoginReplyPacket struct {
	reply int32
}

func (LoginReplyPacket) GetID() uint16 {
	return BanchoLoginReply
}

func (p LoginReplyPacket) GetData() []byte {
	return osubinary.Int32(p.reply)
}

func LoginReply(reply int32) LoginReplyPacket {
	return LoginReplyPacket{reply}
}
