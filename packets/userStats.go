package packets

import "github.com/Mempler/osubinary"

type UserStatsPacket struct {
	UserID          uint32
	Status          byte
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        byte
	BeatmapID       uint32
	RankedScore     uint64
	Accuracy        float32
	PlayCount       uint32
	TotalScore      uint64
	Rank            uint32
	PP              uint16
}

func (UserStatsPacket) GetID() uint16 {
	return BanchoHandleOsuUpdate
}

func (p UserStatsPacket) GetData() []byte {
	return osubinary.Marshal(p)
}
