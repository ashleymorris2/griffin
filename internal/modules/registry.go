package modules

type OutputHandler func(string)

func Register(handler OutputHandler) *ModuleRegistry {
	registry := NewModuleRegistry()

	registry.Register("shell", &ShellModule{
		handler,
	})

	return registry
}
