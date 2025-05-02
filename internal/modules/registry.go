package modules

type OutputHandler func(string)

func Register() *ModuleRegistry {
	registry := NewModuleRegistry()

	registry.Register("shell", &ShellModule{})

	return registry
}
