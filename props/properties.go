package props

import (
	"os"
)

// This type represents a collection of key/value properties, and is
// used for both the env properties and the command line arguments.
type Properties []string

// Returns a copy of the system environment as a props.Properties type
func DefaultEnv() *Properties {
	var defaultEnv Properties = os.Environ()
	return &defaultEnv
}

// Create an empty collection for the properties.
// More properties can be added later
func Empty() *Properties {
	return &Properties{}
}

// Create from a map of key/values
func From(properties map[string]string) *Properties {
	return Empty().AppendAll(properties)
}

// Returns a copy of the properties that will not mutate the original collection
func (props *Properties) Copy() *Properties {
	var copy = *props
	return &copy
}

// Append a key/value property. Flags can be represented with an empty value
//
// E.g.: props.Append("-flag", "")
func (props *Properties) Append(key, value string) *Properties {
	if value == "" {
		*props = append(*props, key)
		return props
	}
	*props = append(*props, key+"="+value)
	return props
}

// Append multiple key/value properties. Flags can be represented with an empty value
//
// E.g.: props.AppendAll(map[string]string{"-arg": "value", "-flag": ""})
func (props *Properties) AppendAll(properties map[string]string) *Properties {
	if value, ok := properties["-C"]; ok {
		// -C must be the first flag
		props.Append("-C", value)
	}
	for key, value := range properties {
		if key == "-C" {
			continue
		}
		props.Append(key, value)
	}
	return props
}

// props1.Merge(props2) will merge props2 properties into props1,
// but without modifying props2
func (props *Properties) Merge(extraProps *Properties) *Properties {
	*props = append(*props, *extraProps...)
	return props
}
