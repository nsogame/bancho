package packets

import "github.com/Mempler/osubinary"

type UserPresencePacket struct {
	UserID      uint32
	Username    string
	Timezone    byte
	Country     byte
	Permissions byte
	Lon         float64
	Lat         float64
	Rank        uint32
}

func (UserPresencePacket) GetID() uint16 {
	return BanchoUserPresence
}

func (p UserPresencePacket) GetData() []byte {
	return osubinary.Marshal(p)
}
