package packets

import "github.com/Mempler/osubinary"

type ProtocolNegotiationPacket struct {
	version uint32
}

func (ProtocolNegotiationPacket) GetID() uint16 {
	return BanchoProtocolNegotiation
}

func (p ProtocolNegotiationPacket) GetData() []byte {
	return osubinary.UInt32(p.version)
}

func ProtocolVersion(version uint32) ProtocolNegotiationPacket {
	return ProtocolNegotiationPacket{version}
}
