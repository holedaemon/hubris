package application

import "github.com/holedaemon/hubris/internal/discord/api/resources"

type Resource struct {
	rc *resources.RestClient
	id string
}

func NewApplicationResource(rc *resources.RestClient, id string) *Resource {
	return &Resource{
		rc: rc,
		id: id,
	}
}
