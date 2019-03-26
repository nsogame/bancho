package bancho

import "git.iptq.io/nso/bancho/packets"

func (bancho *BanchoServer) GetUserPresences() []uint32 {
	return []uint32{}
}

func (bancho *BanchoServer) GetOnlineFriends() []uint32 {
	return []uint32{}
}

func (bancho *BanchoServer) GetUserStats(id int32) packets.UserStatsPacket {
	return packets.UserStatsPacket{
		UserID: uint32(id),
		Rank:   1,
	}
}
