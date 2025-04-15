package core

type StepStage int

const (
	Initial StepStage = iota
	Checking
	Installing
	Done
)

type Dependency struct {
	Name         string
	CheckCommand string
	Status       StepStage
	Found        bool
	Installed    bool
	InstallURL   string
}

func (d *Dependency) Check() (bool, StepStage) {
	d.Found = DependencyExists(d.CheckCommand)
	if d.Found {
		return true, Done
	} else {
		return false, Installing
	}
}

func (d *Dependency) Install() (bool, StepStage) {
	// TODO: add real install logic (e.g., download, run command, etc.)
	return false, Done
}
