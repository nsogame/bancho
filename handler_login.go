package bancho

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/nsogame/bancho/packets"
	"github.com/nsogame/common"
	"github.com/nsogame/common/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	LoginFailed         = -1
	ClientOutdated      = -2
	LoginBanned         = -3
	LoginMultiAcc       = -4
	LoginException      = -5
	RequireSupporter    = -6
	RequireVerification = -8
)

func writeErr(err error, p packets.Packet, w http.ResponseWriter) {
	log.Println(err)
	w.Header().Add("cho-token", "error")
	packets.Write(p, w)
}

func (bancho *BanchoServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErr(err, packets.LoginReply(LoginException), w)
		return
	}

	fmt.Println("body:", body)

	lines := bytes.Split(body, []byte("\n"))
	fmt.Println("lines:", lines)
	if len(lines) < 3 {
		writeErr(err, packets.LoginReply(LoginException), w)
		return
	}

	// attempt to parse the login data now

	// line 1: username
	username := strings.Trim(string(lines[0]), "\n")
	fmt.Printf("username: '%s'\n", username)

	// line 2: md5(password)
	hash := strings.Trim(string(lines[1]), "\n")
	fmt.Printf("password: '%s'\n", hash)

	// line 3: BuildName|TimeZone|HasCity|ClientHash|BlockNonFriends
	// BuildName: str = build name of the client duh5f4dcc3b5aa765d61d8327deb882cf99
	// TimeZone: int = number of hours off from UTC
	// HasCity: bool = whether or not city info is available
	// ClientHash: str = OsuMd5:Adapters:Md5(Adapters):Md5(UniqueId):Md5(UniqueId2):
	// BlockNonFriends: bool = whether or not to block non-friend DMs
	params := strings.Trim(string(lines[2]), "\n")
	fmt.Printf("params: '%s'\n", params)

	// attempt to auth to the db
	var user models.User
	bancho.db.Where("username = ?", strings.ToLower(username)).First(&user)

	if err = bcrypt.CompareHashAndPassword([]byte(user.OsuPassword), []byte(hash)); err != nil {
		fmt.Println("err: ", err)
		writeErr(err, packets.LoginReply(LoginFailed), w)
		return
	}

	// yay the user is authenticated! make some preparations
	choToken, err := uuid.NewRandom()
	if err != nil {
		writeErr(err, packets.LoginReply(LoginException), w)
		return
	}
	client, err := common.NewClient(choToken.String(), user)
	if err != nil {
		writeErr(err, packets.LoginReply(LoginException), w)
		return
	}
	fmt.Println("authenticated!")
	fmt.Println("client:", client)

	// tell the client the good news
	buf := new(bytes.Buffer)
	packets.Write(packets.ProtocolVersion(19), buf)
	packets.Write(packets.LoginReply(int32(user.ID)), buf)
	packets.Write(packets.UserPresencePacket{Username: user.UsernameCase}, buf)
	packets.Write(packets.FriendsList(bancho.GetOnlineFriends()), buf)
	packets.Write(packets.UserPresenceBundle(bancho.GetUserPresences()), buf)
	packets.Write(bancho.GetUserStats(int32(user.ID)), buf)
	packets.Write(packets.LoginPermissions(1), buf)
	packets.Write(packets.SilenceEnd(-1), buf)
	packets.Write(packets.ChannelJoinSuccess("#nso"), buf)
	fmt.Println("buf:", buf.Bytes())

	w.Header().Add("cho-token", choToken.String())
	w.Write(buf.Bytes())
}
