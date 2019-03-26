package packets

import "github.com/Mempler/osubinary"

type LoginPermissionsPacket struct {
	permission int32
}

func (LoginPermissionsPacket) GetID() uint16 {
	return BanchoLoginPermissions
}

func (p LoginPermissionsPacket) GetData() []byte {
	return osubinary.Int32(p.permission)
}

func LoginPermissions(permission int32) LoginPermissionsPacket {
	return LoginPermissionsPacket{permission}
}
