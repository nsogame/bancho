package bancho

const (
	Userperm = 1 << iota
	BAT
	Supporter
	Moderator
	Developer
	Administrator
	TourneyStuff
)

type Client struct {
}
