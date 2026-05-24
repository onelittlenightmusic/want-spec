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
//
// Supported combinations:
//   - Param + As:              top-down — receive a param from upper scope (global or parent) into local param.
//   - CurrentState + As:       bottom-up — push local current state to parent's state (or global state if no parent).
//   - CurrentState + AsGoal:   bottom-up — push local current state to parent's Goal-labeled state via SetGoal.
//   - CurrentState + AsGlobalParam: write local current state directly to a named global parameter.
type ExposeEntry struct {
	Param          string `json:"param,omitempty" yaml:"param,omitempty"`
	CurrentState   string `json:"currentState,omitempty" yaml:"currentState,omitempty"`
	As             string `json:"as,omitempty" yaml:"as,omitempty"`
	AsGoal         string `json:"asGoal,omitempty" yaml:"asGoal,omitempty"`
	AsGlobalParam  string `json:"asGlobalParam,omitempty" yaml:"asGlobalParam,omitempty"`
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

// Want is the canonical DTO representing a single want.
// It contains only the client-mutable fields (Metadata and Spec).
// Runtime-only fields (goroutines, channels, state maps) live in the engine's internal Want type.
type Want struct {
	Metadata Metadata `json:"metadata" yaml:"metadata"`
	Spec     WantSpec `json:"spec" yaml:"spec"`
}

// WantSpec contains the desired state configuration for a want
type WantSpec struct {
	Params              map[string]any       `json:"params" yaml:"params"`
	Exposes             []ExposeEntry        `json:"exposes,omitempty" yaml:"exposes,omitempty"`
	// Imports maps global state keys to internal state keys.
	// Imported fields are read-only: reads transparently return the current global state value;
	// writes are silently blocked. Values are never copied — they are always resolved live.
	Imports             map[string]string    `json:"imports,omitempty" yaml:"imports,omitempty"`
	Using               []map[string]string  `json:"using,omitempty" yaml:"using,omitempty"`
	Recipe              string               `json:"recipe,omitempty" yaml:"recipe,omitempty"`
	StateSubscriptions  []StateSubscription  `json:"stateSubscriptions,omitempty" yaml:"stateSubscriptions,omitempty"`
	NotificationFilters []NotificationFilter `json:"notificationFilters,omitempty" yaml:"notificationFilters,omitempty"`
	Requires            []string             `json:"requires,omitempty" yaml:"requires,omitempty"`
	When                []WhenSpec           `json:"when,omitempty" yaml:"when,omitempty"`
	FinalResultField    string               `json:"finalResultField,omitempty" yaml:"finalResultField,omitempty"`
	ResetOnRestart      *bool                `json:"resetOnRestart,omitempty" yaml:"resetOnRestart,omitempty"`

	// UnknownFields is populated by UnmarshalJSON / UnmarshalYAML when the input
	// contains keys that this version of the spec does not recognise.  It is
	// intentionally omitted from JSON/YAML output so it does not round-trip.
	// Consumers (e.g. mywant engine) should log a warning when this is non-empty.
	UnknownFields []string `json:"-" yaml:"-"`
}
