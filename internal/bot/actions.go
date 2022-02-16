package bot

import (
	"context"
	"math/rand"
	"regexp"
	"time"

	"github.com/holedaemon/hubris/internal/discord/api/resources/guild"
	"github.com/holedaemon/hubris/internal/discord/types"
)

type ActionType int

const (
	ActionTypeTimeout ActionType = 0
)

var (
	regexWaste = regexp.MustCompile(`computer, ((e-?)?waste|ice|drop) this \w+`)
)

var actions = map[*regexp.Regexp]ActionType{
	regexWaste: ActionTypeTimeout,
}

func (b *Bot) TimeoutUser(ctx context.Context, gid, id string) error {
	g := b.api.Guild(gid)

	f := rand.Intn(10)
	t := time.Now().Add(time.Duration(f) * time.Minute)

	p := &guild.ModifyGuildMemberParams{
		CommunicationDisabledUntil: types.TimeToTime(t),
	}

	_, err := g.ModifyGuildMember(ctx, gid, p)
	return err
}
