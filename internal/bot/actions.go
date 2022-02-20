package bot

import (
	"context"
	"math/rand"
	"regexp"

	"github.com/holedaemon/hubris/internal/discord/types"
)

const errCodeMissingPermissions = 50013

var (
	regexWaste = regexp.MustCompile(`(computer|boys), \w+ this \w+`)
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

	var self bool
	if c.Message.Author.ID == c.Message.ReferencedMessage.Author.ID {
		self = true
	}

	err := c.Timeout(ctx)
	if err != nil {
		switch err := err.(type) {
		case *types.Error:
			if err.Code == errCodeMissingPermissions {
				c.Reply(ctx, reactions["yousuck"])
			}

			c.Reply(ctx, "Sorry's boss, da feds gots in da way: \"%s\"", err.Error())
			return
		default:
			c.Reply(ctx, "Sorry's boss, our's cover was blowns, we'll have to try agains")
			return
		}
	}

	if !self {
		c.Reply(ctx, "its dones, boss")
	} else {
		c.Reply(ctx, "https://holedaemon.net/images/snipes.jpg")
	}
}
