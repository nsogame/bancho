package packets

import "github.com/Mempler/osubinary"

type FriendsListPacket struct {
	friends []uint32
}

func (FriendsListPacket) GetID() uint16 {
	return BanchoFriendsList
}

func (p FriendsListPacket) GetData() []byte {
	return osubinary.UIntArray(p.friends)
}

func FriendsList(friends []uint32) FriendsListPacket {
	return FriendsListPacket{friends}
}
