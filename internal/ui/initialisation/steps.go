package initialisation

// Step identifiers (string constants used as keys for ordering, status mapping, etc.)
const (
	stepPrepareEnv    = "prepare-env"           // Set up the environment (e.g., ensure required dirs exist)
	stepCreateExample = "create-example-config" //Create an example config file (disable with --example=false
)

// stepOrder defines the order in which setup steps should be executed and displayed in the UI.
// The order here controls both the execution sequence and how steps appear in the terminal output.
var stepOrder = []string{
	stepPrepareEnv,
	stepCreateExample,
}
