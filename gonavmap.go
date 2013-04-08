package gonavmap

import (
	"fmt"
	"strings"
)

func Get(m map[string]interface{}, path string) map[string]interface{} {
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return nil
	}
	return navmap(m, parts)
}

func Set(m map[string]interface{}, path string, value interface{}) error {
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return fmt.Errorf("Path provided is too short")
	}

	parent := navmap(m, parts[:len(parts)-1])

	if parent != nil {
		parent[parts[len(parts)-1]] = value
	} else {
		return fmt.Errorf("Invalid path provided")
	}

	return nil
}

func Value(m map[string]interface{}, path string) interface{} {
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return nil
	}

	parent := navmap(m, parts[:len(parts)-1])

	if parent != nil {
		if v, ok := parent[parts[len(parts)-1]]; ok {
			return v
		}
	}

	return nil
}

func navmap(m map[string]interface{}, path []string) map[string]interface{} {
	if len(path) == 0 {
		return m
	}

	if element, ok := m[path[0]]; ok {
		if m, ok := element.(map[string]interface{}); ok {
			return navmap(m, path[1:])
		} else {
			return nil
		}
	}

	return nil
}
