package initialisation

// Step identifiers (string constants used as keys for ordering, status mapping, etc.)
const (
	stepPrepareEnv = "prepare-env" // Set up the environment (e.g., ensure required dirs exist)
)

// stepOrder defines the order in which setup steps should be executed and displayed in the UI.
// The order here controls both the execution sequence and how steps appear in the terminal output.
var stepOrder = []string{
	stepPrepareEnv,
}

// stepLabels maps internal step IDs to human-friendly display names.
// These are shown in the terminal UI, and are separate from the internal identifiers to keep things flexible.
// You can safely change these without affecting program logic.
var stepLabels = map[string]string{
	stepPrepareEnv: "Preparing environment",
}
