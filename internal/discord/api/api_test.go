package api

import (
	"context"
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
