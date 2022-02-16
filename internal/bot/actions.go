package bot

import (
	"context"
	"errors"
	"math/rand"
	"regexp"

	"github.com/holedaemon/hubris/internal/discord/types"
)

var (
	regexWaste = regexp.MustCompile(`(computer|boys), ((e-?)?waste|ice|drop) this \w+`)
)

var actions = map[*regexp.Regexp]func(context.Context, *Context){
	regexWaste: actionWaste,
}

func actionWaste(ctx context.Context, c *Context) {
	s := rand.Intn(99)
	if s == 0 {
		c.Reply(ctx, "why mafia isn't a fucking aesthetic: a thread")
		return
	}

	if c.Message.Type != types.MessageTypeReply {
		c.Reply(ctx, "Who's you'se wanting me to hit's, boss?")
		return
	}

	err := c.Timeout(ctx)
	if err != nil {
		switch err := err.(type) {
		case *types.Error:
			c.Reply(ctx, "Sorry's boss, da feds gots in da way: \"%s\"", err.Error())
			return
		default:
			if errors.Is(err, errNotReply) {
				c.Reply(ctx, "Sorry's boss, I am legally obligated to never lay's a finger on youse.")
				return
			}

			c.Reply(ctx, "Sorry's boss, our's cover was blowns, we'll have to try agains")
			return
		}
	}

	c.Reply(ctx, "its dones, boss")
}
