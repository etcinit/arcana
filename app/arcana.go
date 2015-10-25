package app

import (
	"github.com/etcinit/arcana/app/workflows"
	"github.com/jacobstr/confer"
)

// Arcana is the root node of the DI graph
type Arcana struct {
	Config     *confer.Config               `inject:""`
	Diffusion  *workflows.DiffusionService  `inject:""`
	Maniphest  *workflows.ManiphestService  `inject:""`
	Repository *workflows.RepositoryService `inject:""`
}

// Boot boots the upper part of the application.
func (p *Arcana) Boot() {
	//
}
