package workflows

import (
	"github.com/codegangsta/cli"
	"github.com/etcinit/gonduit"
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/phabulous/app/factories"
)

// RepositoryService provides commands for interacting with the repositories
// application.
type RepositoryService struct {
	GonduitFactory *factories.GonduitFactory `inject:""`
}

// QueryByCallsign queries for repositories by their callsigns.
func (d *RepositoryService) QueryByCallsign(c *cli.Context) {
	spewOrPanic(d.GonduitFactory, func(client *gonduit.Conn) (interface{}, error) {
		return client.RepositoryQuery(requests.RepositoryQueryRequest{
			Callsigns: []string(c.Args()),
			Order:     "newest",
		})
	})
}
