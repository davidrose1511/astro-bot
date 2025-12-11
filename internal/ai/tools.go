package ai

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

// GetTools returns the OpenAI-formatted tool definitions
// Notice: return type is []openai.Tool (NOT genai.Tool)
func GetTools() []openai.Tool {
	return []openai.Tool{
		{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "get_kundali_details",
				Description: "Extract birth details. Call this when user gives Name, DOB, TOB, and City.",
				Parameters: jsonschema.Definition{
					Type: jsonschema.Object,
					Properties: map[string]jsonschema.Definition{
						"name": {
							Type:        jsonschema.String,
							Description: "Name of user",
						},
						"dob": {
							Type:        jsonschema.String,
							Description: "Date of Birth (YYYY-MM-DD)",
						},
						"tob": {
							Type:        jsonschema.String,
							Description: "Time of Birth (HH:MM)",
						},
						"city": {
							Type:        jsonschema.String,
							Description: "City of birth",
						},
					},
					Required: []string{"dob", "tob", "city"},
				},
			},
		},
	}
}