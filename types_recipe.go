package want_spec

// RecipeWant represents a want in recipe format (aligned with Want structure)
type RecipeWant struct {
	Metadata Metadata `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec     WantSpec `yaml:"spec,omitempty" json:"spec,omitempty"`

	// Legacy flattened fields for backward compatibility
	Name        string              `yaml:"name,omitempty" json:"name,omitempty"`
	Type        string              `yaml:"type,omitempty" json:"type,omitempty"`
	Labels      map[string]string   `yaml:"labels,omitempty" json:"labels,omitempty"`
	Params      map[string]any      `yaml:"params,omitempty" json:"params,omitempty"`
	Using       []map[string]string `yaml:"using,omitempty" json:"using,omitempty"`
	Requires    []string            `yaml:"requires,omitempty" json:"requires,omitempty"`
	RecipeAgent bool                `yaml:"recipeAgent,omitempty" json:"recipeAgent,omitempty"`
}

// GenericRecipe represents the top-level recipe structure
type GenericRecipe struct {
	Recipe RecipeContent `yaml:"recipe" json:"recipe"`
}

// RecipeResult defines how to compute results from recipe execution
type RecipeResult []RecipeResultSpec

// RecipeResultSpec specifies which want and state field to use for result computation
type RecipeResultSpec struct {
	WantName    string `yaml:"want_name" json:"want_name"`
	StateField  string `yaml:"state_field,omitempty" json:"state_field,omitempty"`
	StatName    string `yaml:"stat_name,omitempty" json:"stat_name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// RecipeExample represents example deployment configuration for one-click deployment
type RecipeExample struct {
	Wants []RecipeWant `yaml:"wants,omitempty" json:"wants,omitempty"`
}

// RecipeExampleDef defines a named example usage of a recipe with pre-filled parameters
type RecipeExampleDef struct {
	Name             string         `json:"name" yaml:"name"`
	Description      string         `json:"description" yaml:"description"`
	Params           map[string]any `json:"params" yaml:"params"`
	ExpectedBehavior string         `json:"expectedBehavior,omitempty" yaml:"expectedBehavior,omitempty"`
}

// RecipeContent contains the actual recipe data
type RecipeContent struct {
	Metadata              GenericRecipeMetadata `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Wants                 []RecipeWant          `yaml:"wants,omitempty" json:"wants,omitempty"`
	Parameters            map[string]any        `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	ParameterDescriptions map[string]string     `yaml:"parameter_descriptions,omitempty" json:"parameter_descriptions,omitempty"`
	Result                *RecipeResult         `yaml:"result,omitempty" json:"result,omitempty"`
	Example               *RecipeExample        `yaml:"example,omitempty" json:"example,omitempty"`
	Examples              []RecipeExampleDef    `yaml:"examples,omitempty" json:"examples,omitempty"`
	State                 []StateDef            `yaml:"state,omitempty" json:"state,omitempty"`
	FinalResultField      string                `yaml:"finalResultField,omitempty" yaml:"finalResultField,omitempty"`
}

// GenericRecipeMetadata contains recipe information
type GenericRecipeMetadata struct {
	ID          string `yaml:"id,omitempty" json:"id,omitempty"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Version     string `yaml:"version" json:"version"`
	Type        string `yaml:"type,omitempty" json:"type,omitempty"`
	CustomType  string `yaml:"custom_type,omitempty" json:"custom_type,omitempty"`
	Category    string `yaml:"category,omitempty" json:"category,omitempty"`
}
