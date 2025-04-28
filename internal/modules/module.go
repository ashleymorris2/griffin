package modules

import "github.com/ashleymorris2/booty/internal/models"

type Module interface {
	Run(t models.Task) error
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
