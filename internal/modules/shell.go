package modules

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ashleymorris2/booty/internal/models"
	"os/exec"
)

type ShellModule struct {
	OnOutput func(line string)
}

func (m *ShellModule) Name() string {
	return "shell"
}

func (m *ShellModule) Run(t models.Task) error {
	script, ok := t.With["script"].(string)
	if !ok || script == "" {
		return fmt.Errorf("missing 'script' in task: %s", t.Label)
	}

	cmd := exec.CommandContext(context.Background(), "sh", "-c", script)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if m.OnOutput != nil {
		if out := stdout.String(); out != "" {
			m.OnOutput(out)
		}
		if errOut := stderr.String(); errOut != "" {
			m.OnOutput(errOut)
		}
	}

	return err
}
