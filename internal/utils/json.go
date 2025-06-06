package utils

import "fmt"

func JsonExtractInner(raw map[string]any, properties ...string) (map[string]any, error) {
	if len(properties) == 0 {
		return nil, fmt.Errorf("Must pass at least 1 property for extraction")
	}

	for _, prop := range properties {
		data, ok := raw[prop].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("Property %s does not exist in data", prop)
		}

		raw = data
	}

	return raw, nil
}
