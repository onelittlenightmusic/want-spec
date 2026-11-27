package want_spec

// Character is someone on the board — a person, played from a device, or one
// acting as a robot (label robot=true): moved by the backend, answered by the
// robot's model with its tools. What they see is theirs (their Story, what
// they remember hearing), and so is when it is worth telling: a character has
// timings, as a want type does, read off the character's own state.
//
//	apiVersion: mywant.io/v1
//	kind: Character
//	metadata:
//	  id: chr-a8092b7d
//	  name: aaa
//	  labels: {robot: "true"}
//	spec:
//	  avatar: 🦊
//	  color: "#0ea5e9"
//	  timings:
//	    - name: errand-different
//	      description: お使いで想定と違った
//	      if: errand_different
//	      notify: true
//
//	GET  /api/v1/characters/{id}?format=yaml   Character
//	POST /api/v1/characters/apply              Character (YAML or JSON) → the character as kept
//
// See spec/character-spec.yaml.
type Character struct {
	APIVersion string            `yaml:"apiVersion" json:"apiVersion"`
	Kind       string            `yaml:"kind" json:"kind"`
	Metadata   CharacterMetadata `yaml:"metadata" json:"metadata"`
	Spec       CharacterSpec     `yaml:"spec" json:"spec"`
}

// CharacterKind is a Character document's kind.
const CharacterKind = "Character"

// CharacterMetadata is who a character is: their identity (ID, made once),
// what they are called (Name) and their labels.
type CharacterMetadata struct {
	ID     string            `yaml:"id,omitempty" json:"id,omitempty"`
	Name   string            `yaml:"name" json:"name"`
	Labels map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

// CharacterSpec is how a character looks, and when what happens to them is
// worth telling.
type CharacterSpec struct {
	// Avatar is the emoji they are drawn as.
	Avatar string `yaml:"avatar,omitempty" json:"avatar,omitempty"`
	// Color is theirs, #rrggbb.
	Color string `yaml:"color,omitempty" json:"color,omitempty"`
	// Shape is the outline they stand in ("circle", "star", …).
	Shape string `yaml:"shape,omitempty" json:"shape,omitempty"`
	// Timings: the moments worth telling of, as a want type's are — If,
	// OnChange, OnIncrease, OnDecrease read off the character's state
	// (CharacterState*), Murmur (said in their bubble) and Notify (a notice
	// from them). One named as a standard timing (CharacterStandardTimings)
	// takes its place; the rest of the standard ones hold as they are.
	Timings []TimingDef `yaml:"timings,omitempty" json:"timings,omitempty"`
}

// The label a character acting as a robot carries ("true").
const CharacterLabelRobot = "robot"

// A character's state: the fields a character's timings read. Each is read
// when something happens to them — they hear something, come back from an
// errand, move.
const (
	// CharacterStateMood is how they feel: joy, eager, calm, waiting,
	// worried, sad — the mood of what moved them most lately.
	CharacterStateMood = "mood"
	// CharacterStateIntensity is how strongly, 0 to 1.
	CharacterStateIntensity = "intensity"
	// CharacterStateMoodUneasy is true when they are worried or sad.
	CharacterStateMoodUneasy = "mood_uneasy"
	// CharacterStatePlace is the name of what they stand on.
	CharacterStatePlace = "place"
	// CharacterStateHeard is the newest thing they remember hearing.
	CharacterStateHeard = "heard"
	// CharacterStateErrandVerdict is how the errand they are back from went:
	// same, different, unknown, or first (nothing was expected). Read only
	// as they come back.
	CharacterStateErrandVerdict = "errand_verdict"
	// CharacterStateErrandDifferent is true as they come back from an errand
	// that found something other than they expected.
	CharacterStateErrandDifferent = "errand_different"
	// CharacterStateErrandReport is what they said of it.
	CharacterStateErrandReport = "errand_report"
)

// CharacterStandardTimings are every character's unless their spec names one
// the same: an errand that found other than expected, told; a mood turned
// uneasy, said in their bubble.
func CharacterStandardTimings() []TimingDef {
	return []TimingDef{
		{
			Name: "errand-different", Description: "お使いで想定と違った", Icon: "map-pin-off",
			If: CharacterStateErrandDifferent, Notify: true, Irregular: true,
		},
		{
			Name: "uneasy", Description: "気になることができた", Icon: "cloud-rain",
			If: CharacterStateMoodUneasy, OnChange: CharacterStateMood, Murmur: true,
		},
	}
}

// CharacterTimings are a character's timings: the standard ones, each
// replaced by one of theirs of the same name, then the rest of theirs.
func CharacterTimings(own []TimingDef) []TimingDef {
	named := map[string]TimingDef{}
	for _, t := range own {
		named[t.Name] = t
	}
	var out []TimingDef
	for _, t := range CharacterStandardTimings() {
		if mine, ok := named[t.Name]; ok {
			out = append(out, mine)
			delete(named, t.Name)
			continue
		}
		out = append(out, t)
	}
	for _, t := range own {
		if _, ok := named[t.Name]; ok {
			out = append(out, t)
		}
	}
	return out
}
