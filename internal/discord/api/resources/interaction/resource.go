package interaction

import "github.com/holedaemon/hubris/internal/discord/api/resources"

type Resource struct {
	appID string
	id    string
	token string

	rc *resources.RestClient
}

func NewResource(appID, id, token string, rc *resources.RestClient) *Resource {
	return &Resource{
		appID: appID,
		id:    id,
		token: token,
		rc:    rc,
	}
}
