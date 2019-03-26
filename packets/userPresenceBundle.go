package packets

import "github.com/Mempler/osubinary"

type UserPresenceBundlePacket struct {
	users []uint32
}

func (UserPresenceBundlePacket) GetID() uint16 {
	return BanchoUserPresenceBundle
}

func (p UserPresenceBundlePacket) GetData() []byte {
	return osubinary.UIntArray(p.users)
}

func UserPresenceBundle(users []uint32) UserPresenceBundlePacket {
	return UserPresenceBundlePacket{users}
}
