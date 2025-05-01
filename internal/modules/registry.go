package modules

type OutputHandler func(string)

func RegisterModules() *ModuleRegistry {
	registry := NewModuleRegistry()

	registry.Register("shell", &ShellModule{})

	return registry
}
