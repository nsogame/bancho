package packets

import "github.com/Mempler/osubinary"

type SilenceEndPacket struct {
	timeout int32
}

func (SilenceEndPacket) GetID() uint16 {
	return BanchoSilenceEnd
}

func (p SilenceEndPacket) GetData() []byte {
	return osubinary.Int32(p.timeout)
}

func SilenceEnd(timeout int32) SilenceEndPacket {
	return SilenceEndPacket{timeout}
}
