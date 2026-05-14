package want_spec

// OwnerReference represents a reference to an owner object
type OwnerReference struct {
	APIVersion         string `json:"apiVersion" yaml:"apiVersion"`
	Kind               string `json:"kind" yaml:"kind"`
	Name               string `json:"name" yaml:"name"`
	ID                 string `json:"id" yaml:"id"`
	Controller         bool   `json:"controller,omitempty" yaml:"controller,omitempty"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion,omitempty" yaml:"blockOwnerDeletion,omitempty"`
}

// CorrelationEntry represents the correlation relationship between two Wants.
type CorrelationEntry struct {
	WantID string   `json:"wantID" yaml:"wantID"`
	Labels []string `json:"labels" yaml:"labels"`
	Rate   int      `json:"rate"   yaml:"rate"`
}

// Metadata contains want identification and classification info
type Metadata struct {
	ID              string             `json:"id,omitempty" yaml:"id,omitempty"`
	Name            string             `json:"name" yaml:"name"`
	Type            string             `json:"type" yaml:"type"`
	Labels          map[string]string  `json:"labels" yaml:"labels"`
	OwnerReferences []OwnerReference   `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	UpdatedAt       int64              `json:"updatedAt" yaml:"-"`
	IsSystemWant    bool               `json:"isSystemWant,omitempty" yaml:"isSystemWant,omitempty"`
	OrderKey        string             `json:"orderKey,omitempty" yaml:"orderKey,omitempty"`
	Correlation     []CorrelationEntry `json:"correlation,omitempty" yaml:"correlation,omitempty"`
	Series          string             `json:"series,omitempty" yaml:"series,omitempty"`
	Version         int                `json:"version,omitempty" yaml:"version,omitempty"`
}

// StateSubscription defines what state changes to monitor
type StateSubscription struct {
	WantName   string   `json:"wantName" yaml:"wantName"`
	StateKeys  []string `json:"stateKeys,omitempty" yaml:"stateKeys,omitempty"`
	Conditions []string `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	BufferSize int      `json:"bufferSize,omitempty" yaml:"bufferSize,omitempty"`
}

// ParamEntry represents a single parameter entry in array format
type ParamEntry struct {
	Key   string `json:"key" yaml:"key"`
	Value any    `json:"value" yaml:"value"`
}

// ExposeEntry declares a parameter or state exposure between scope levels.
type ExposeEntry struct {
	Param        string `json:"param,omitempty" yaml:"param,omitempty"`
	CurrentState string `json:"currentState,omitempty" yaml:"currentState,omitempty"`
	As           string `json:"as" yaml:"as"`
}

// NotificationFilter allows filtering received notifications
type NotificationFilter struct {
	SourcePattern string   `json:"sourcePattern" yaml:"sourcePattern"`
	StateKeys     []string `json:"stateKeys,omitempty" yaml:"stateKeys,omitempty"`
	ValuePattern  string   `json:"valuePattern,omitempty" yaml:"valuePattern,omitempty"`
}

// ParamRef allows a spec.params value to reference a global parameter by name,
// using the same fromGlobalParam convention as WhenSpec.
// Example YAML: budget: {fromGlobalParam: global_budget}
type ParamRef struct {
	FromGlobalParam string `json:"fromGlobalParam" yaml:"fromGlobalParam"`
}

// WhenSpec defines a scheduled execution time for a Want
type WhenSpec struct {
	At              string `json:"at,omitempty" yaml:"at,omitempty"`
	Every           string `json:"every,omitempty" yaml:"every,omitempty"`
	FromGlobalParam string `json:"fromGlobalParam,omitempty" yaml:"fromGlobalParam,omitempty"`
}

// WantSpec contains the desired state configuration for a want
type WantSpec struct {
	Params              map[string]any       `json:"params" yaml:"params"`
	Exposes             []ExposeEntry        `json:"exposes,omitempty" yaml:"exposes,omitempty"`
	Using               []map[string]string  `json:"using,omitempty" yaml:"using,omitempty"`
	Recipe              string               `json:"recipe,omitempty" yaml:"recipe,omitempty"`
	StateSubscriptions  []StateSubscription  `json:"stateSubscriptions,omitempty" yaml:"stateSubscriptions,omitempty"`
	NotificationFilters []NotificationFilter `json:"notificationFilters,omitempty" yaml:"notificationFilters,omitempty"`
	Requires            []string             `json:"requires,omitempty" yaml:"requires,omitempty"`
	When                []WhenSpec           `json:"when,omitempty" yaml:"when,omitempty"`
	FinalResultField    string               `json:"finalResultField,omitempty" yaml:"finalResultField,omitempty"`
	ResetOnRestart      *bool                `json:"resetOnRestart,omitempty" yaml:"resetOnRestart,omitempty"`
}
