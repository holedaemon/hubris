package api

import (
	"context"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

const bogusToken = "hjgdhgkjsdgkjshgkdsjfhgfdkjghfkjlol"

func TestGetGateway(t *testing.T) {
	c, err := New(bogusToken)
	assert.NilError(t, err, "creating client")

	url, err := c.GetGateway(context.Background())
	assert.NilError(t, err, "getting gateway URL")

	t.Log(url)
}

func TestGetChannel(t *testing.T) {
	token := os.Getenv("TOKEN")
	assert.Assert(t, token != "", "$TOKEN is blank")

	chID := os.Getenv("CHANNEL_ID")
	assert.Assert(t, chID != "", "$CHANNEL_ID is blank")

	c, err := New(token)
	assert.NilError(t, err, "creating client")

	cr := c.Channel(chID)
	ch, err := cr.Get(context.Background())
	assert.NilError(t, err, "getting channel")

	t.Log(ch)
}
