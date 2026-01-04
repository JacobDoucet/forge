package types

import (
	"strings"
)

type Actor struct {
	Id             string          `yaml:"id"`
	Language       string          `yaml:"language"`
	Username       string          `yaml:"username"`
	Name           string          `yaml:"name"`
	AdminName      string          `yaml:"adminName"`
	LanguageParts  []ActorNamePart `yaml:"-"`
	UsernameParts  []ActorNamePart `yaml:"-"`
	NameParts      []ActorNamePart `yaml:"-"`
	AdminNameParts []ActorNamePart `yaml:"-"`
}

type ActorNamePart struct {
	Field *Field
	Text  string
}

func ParseActorNamePattern(obj *Object, namePattern string) []ActorNamePart {
	if namePattern == "" {
		return []ActorNamePart{
			{Text: "Unnamed Actor"},
		}
	}

	var result []ActorNamePart

	// Split the pattern into segments based on curly braces
	segments := strings.Split(namePattern, "{")

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		// Check if the segment contains a field placeholder
		parts := strings.SplitN(segment, "}", 2)
		if len(parts) > 1 {
			// Extract field name and remaining text
			fieldName := parts[0]
			remainingText := parts[1]

			// Try to retrieve the field from the object
			if field, found := obj.GetField(fieldName); found {
				result = append(result, ActorNamePart{Field: &field})
			} else {
				// If the field is not found, treat the segment as plain text
				result = append(result, ActorNamePart{Text: "{" + segment + "}"})
			}

			// Add any remaining text as a plain text part
			if remainingText != "" {
				result = append(result, ActorNamePart{Text: remainingText})
			}
		} else {
			// If no field placeholder, treat the segment as plain text
			result = append(result, ActorNamePart{Text: segment})
		}
	}

	return result
}
