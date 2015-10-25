package workflows

import (
	"github.com/codegangsta/cli"
	"github.com/etcinit/gonduit"
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/phabulous/app/factories"
)

// DiffusionService provides commands for interacting with the Diffusion
// application.
type DiffusionService struct {
	GonduitFactory *factories.GonduitFactory `inject:""`
}

// QueryCommitsByName queries for commits by their name.
func (d *DiffusionService) QueryCommitsByName(c *cli.Context) {
	spewOrPanic(d.GonduitFactory, func(client *gonduit.Conn) (interface{}, error) {
		return client.DiffusionQueryCommits(requests.DiffusionQueryCommitsRequest{
			Names: []string(c.Args()),
		})
	})
}
