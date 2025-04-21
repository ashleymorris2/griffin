package modules

type Module struct {
	ID        string            `yaml:"id,omitempty"`
	Label     string            `yaml:"label"`
	Uses      string            `yaml:"uses"`
	With      map[string]string `yaml:"with"`
	Platforms []string          `yaml:"platforms,omitempty"`
	DependsOn []string          `yaml:"depends_on,omitempty"`
}
