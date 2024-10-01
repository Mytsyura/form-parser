package main

// Question represents a single form question
type Question struct {
	Label        string                `json:"label" yaml:"label"`
	Type         string                `json:"type" yaml:"type"` // "text", "number", etc.
	DefaultValue string                `json:"default_value,omitempty" yaml:"default_value,omitempty"`
	Options      []string              `json:"options,omitempty" yaml:"options,omitempty"` // For select-type questions (e.g., pizza/burger)
	Next         map[string][]Question `json:"next,omitempty" yaml:"next,omitempty"`       // Follow-up questions based on answer
}

// Form represents the structure of the form
type Form struct {
	Title     string     `json:"title" yaml:"title"`
	Questions []Question `json:"questions" yaml:"questions"`
}
