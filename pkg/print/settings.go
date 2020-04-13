package print

// Settings represents all settings
type Settings struct {
	// EscapeCharacters escapes special characters (such as _ * in Markdown and > < in JSON) (default: true)
	// scope: Markdown
	EscapeCharacters bool

	// EscapePipe escapes pipe character in Markdown (default: true)
	// scope: Markdown
	EscapePipe bool

	// IndentLevel control the indentation of AsciiDoc and Markdown headers [available: 1, 2, 3, 4, 5] (default: 2)
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues ailrghaekrgj
	// scope: Global
	OutputValues bool

	// ShowColor print "colorized" version of result in the terminal (default: true)
	// scope: Pretty
	ShowColor bool

	// ShowHeader show "Header" module information (default: true)
	// scope: Global
	ShowHeader bool

	// ShowInputs show "Inputs" information (default: true)
	// scope: Global
	ShowInputs bool

	// ShowOutputs show "Outputs" information (default: true)
	// scope: Global
	ShowOutputs bool

	// ShowProviders show "Providers" information (default: true)
	// scope: Global
	ShowProviders bool

	// ShowRequired show "Required" column when generating Markdown (default: true)
	// scope: Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" column when generating Markdown (default: true)
	// scope: Markdown
	ShowSensitivity bool

	// ShowRequirements show "Requirements" section (default: true)
	// scope: Global
	ShowRequirements bool

	// SortByName sorted rendering of inputs and outputs (default: true)
	// scope: Global
	SortByName bool

	// SortByRequired sort items (inputs, providers) by name and prints required ones first (default: false)
	// scope: Global
	SortByRequired bool

	// SortByType sort items (inputs, outputs) by type alphabetically (default: false)
	// scope: Global
	SortByType bool
}

//NewSettings returns new instance of Settings
func NewSettings() *Settings {
	return &Settings{
		EscapeCharacters: true,
		EscapePipe:       true,
		IndentLevel:      2,
		OutputValues:     false,
		ShowColor:        true,
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequired:     true,
		ShowSensitivity:  true,
		ShowRequirements: true,
		SortByName:       true,
		SortByRequired:   false,
		SortByType:       false,
	}
}
