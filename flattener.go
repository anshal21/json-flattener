package flattener

import (
	"encoding/json"
	"fmt"
)

// Separator ...
type Separator string

func (s Separator) String() string {
	return string(s)
}

// Separators ...
const (
	DotSeparator Separator = "."
)

type flattenerConfig struct {
	ignoreArray bool
	depth       *int
	prefixes    map[string]bool
}

// Option ...
type Option func(f *flattenerConfig)

// IgnoreArray option, if enabled, ignores arrays while flattening
func IgnoreArray() Option {
	return func(f *flattenerConfig) {
		f.ignoreArray = true
	}
}

// WithDepth option, if provided, limits the flattening to the specified depth
func WithDepth(depth int) Option {
	return func(f *flattenerConfig) {
		f.depth = &depth
	}
}

// FlattenJSON flattens the provided JSON
// The flattening can be customised by providing flattening Options
func FlattenJSON(JSONStr string, separator Separator, options ...Option) (string, error) {
	data := make(map[string]interface{})
	finalMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(JSONStr), &data); err != nil {
		return "", err
	}

	config := &flattenerConfig{}
	for _, option := range options {
		option(config)
	}

	if err := flatten(data, "", separator, config, finalMap, 0); err != nil {
		return "", err
	}

	return mustToJSONStr(finalMap), nil
}

// flatten ....
func flatten(data interface{}, prefix string, separator Separator, config *flattenerConfig, finalMap map[string]interface{}, depth int) error {

	if config.depth != nil && depth == *config.depth {
		finalMap[prefix] = data
		return nil
	}

	switch data.(type) {
	case map[string]interface{}:
		for key, val := range data.(map[string]interface{}) {
			if err := flatten(val, appendToPrefix(prefix, key, separator), separator, config, finalMap, depth+1); err != nil {
				return err
			}
		}
	case []interface{}:
		if config.ignoreArray {
			finalMap[prefix] = data
			return nil
		}
		for index, val := range data.([]interface{}) {
			if err := flatten(val, appendToPrefix(prefix, fmt.Sprintf("%v", index), separator), separator, config, finalMap, depth+1); err != nil {
				return err
			}
		}

	default:
		finalMap[prefix] = data
	}

	return nil
}

func appendToPrefix(prefix string, key string, separator Separator) string {
	if prefix == "" {
		return key
	}
	return fmt.Sprintf("%v%v%v", prefix, separator.String(), key)
}

func mustToJSONStr(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(jsonData)
}
