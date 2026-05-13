package want_spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

// GetParam returns the value for the given key from Params map
func (s *WantSpec) GetParam(key string) (any, bool) {
	v, ok := s.Params[key]
	return v, ok
}

// SetParam sets a param value, initializing the map if nil
func (s *WantSpec) SetParam(key string, val any) {
	if s.Params == nil {
		s.Params = make(map[string]any)
	}
	s.Params[key] = val
}

// HasParam returns true if the key exists in Params
func (s *WantSpec) HasParam(key string) bool {
	_, ok := s.Params[key]
	return ok
}

// ParamsAsMap returns the Params map directly
func (s *WantSpec) ParamsAsMap() map[string]any {
	return s.Params
}

// SetParamsFromMap replaces Params with the given map
func (s *WantSpec) SetParamsFromMap(m map[string]any) {
	s.Params = m
}

// UnmarshalJSON implements json.Unmarshaler to support both old map format and new array format.
func (s *WantSpec) UnmarshalJSON(data []byte) error {
	type WantSpecNoParams struct {
		Exposes             []ExposeEntry        `json:"exposes,omitempty"`
		Using               []map[string]string  `json:"using,omitempty"`
		Recipe              string               `json:"recipe,omitempty"`
		StateSubscriptions  []StateSubscription  `json:"stateSubscriptions,omitempty"`
		NotificationFilters []NotificationFilter `json:"notificationFilters,omitempty"`
		Requires            []string             `json:"requires,omitempty"`
		When                []WhenSpec           `json:"when,omitempty"`
		FinalResultField    string               `json:"finalResultField,omitempty"`
		ResetOnRestart      *bool                `json:"resetOnRestart,omitempty"`
	}
	var raw struct {
		WantSpecNoParams
		Params json.RawMessage `json:"params"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.Exposes = raw.Exposes
	s.Using = raw.Using
	s.Recipe = raw.Recipe
	s.StateSubscriptions = raw.StateSubscriptions
	s.NotificationFilters = raw.NotificationFilters
	s.Requires = raw.Requires
	s.When = raw.When
	s.FinalResultField = raw.FinalResultField
	s.ResetOnRestart = raw.ResetOnRestart

	if raw.Params == nil {
		return nil
	}

	trimmed := bytes.TrimSpace(raw.Params)
	if len(trimmed) > 0 && trimmed[0] == '[' {
		var entries []ParamEntry
		if err := json.Unmarshal(raw.Params, &entries); err != nil {
			return err
		}
		s.Params = make(map[string]any, len(entries))
		for _, e := range entries {
			s.Params[e.Key] = e.Value
		}
	} else {
		var m map[string]any
		if err := json.Unmarshal(raw.Params, &m); err != nil {
			return err
		}
		s.Params = m
	}
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler
func (s *WantSpec) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("WantSpec must be a YAML mapping")
	}

	type restSpec struct {
		Exposes             []ExposeEntry        `yaml:"exposes,omitempty"`
		Using               []map[string]string  `yaml:"using,omitempty"`
		Recipe              string               `yaml:"recipe,omitempty"`
		StateSubscriptions  []StateSubscription  `yaml:"stateSubscriptions,omitempty"`
		NotificationFilters []NotificationFilter `yaml:"notificationFilters,omitempty"`
		Requires            []string             `yaml:"requires,omitempty"`
		When                []WhenSpec           `yaml:"when,omitempty"`
		FinalResultField    string               `yaml:"finalResultField,omitempty"`
		ResetOnRestart      *bool                `yaml:"resetOnRestart,omitempty"`
	}

	var paramsNode *yaml.Node
	var rest restSpec

	for i := 0; i < len(value.Content); i += 2 {
		key := value.Content[i].Value
		val := value.Content[i+1]

		if key == "params" {
			paramsNode = val
			continue
		}

		tempMap := map[string]*yaml.Node{key: val}
		tempData, _ := yaml.Marshal(tempMap)
		yaml.Unmarshal(tempData, &rest)
	}

	s.Exposes = rest.Exposes
	s.Using = rest.Using
	s.Recipe = rest.Recipe
	s.StateSubscriptions = rest.StateSubscriptions
	s.NotificationFilters = rest.NotificationFilters
	s.Requires = rest.Requires
	s.When = rest.When
	s.FinalResultField = rest.FinalResultField
	s.ResetOnRestart = rest.ResetOnRestart

	if paramsNode != nil {
		if paramsNode.Kind == yaml.SequenceNode {
			var entries []ParamEntry
			if err := paramsNode.Decode(&entries); err != nil {
				return err
			}
			s.Params = make(map[string]any, len(entries))
			for _, e := range entries {
				s.Params[e.Key] = e.Value
			}
		} else {
			var m map[string]any
			if err := paramsNode.Decode(&m); err != nil {
				return err
			}
			s.Params = m
		}
	}

	return nil
}
