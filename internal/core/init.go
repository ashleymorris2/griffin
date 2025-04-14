package core

import (
	"os/exec"
)

func DependencyExists(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return true
}
