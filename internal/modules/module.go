package modules

import "github.com/ashleymorris2/booty/internal/core/blueprint"

type Module interface {
	Run(task blueprint.Task) error
}
