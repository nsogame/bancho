package packets

import (
	"encoding/binary"
	"io"
)

type Packet interface {
	GetID() uint16
	GetData() []byte
}

func Write(p Packet, w io.Writer) {
	data := p.GetData()
	binary.Write(w, binary.LittleEndian, p.GetID())
	binary.Write(w, binary.LittleEndian, byte(0))
	binary.Write(w, binary.LittleEndian, uint32(len(data)))
	binary.Write(w, binary.LittleEndian, data)
}
