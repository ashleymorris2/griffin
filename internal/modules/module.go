package modules

import "github.com/ashleymorris2/booty/internal/models"

type Module interface {
	Name() string            // Friendly name of this module e.g. "docker, git, ssh"
	Run(t models.Task) error // The specific implementation of this module
}

type ModuleRegistry struct {
	modules map[string]Module
}

func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules: make(map[string]Module),
	}
}

func (r *ModuleRegistry) Register(name string, module Module) {
	r.modules[name] = module
}

func (r *ModuleRegistry) Get(name string) (Module, bool) {
	module, ok := r.modules[name]
	return module, ok
}
