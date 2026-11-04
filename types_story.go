package want_spec

// A character's story: what the person playing it did on the board — where
// they went from and to, what they opened, what they asked and what they were
// answered — kept by the server from what it is sent anyway (the cursor's
// moves, the GUI's state, the turns and acts), so a chat's answers can be told
// good or bad by what the person did next.
//
//	GET /api/v1/characters/{id}/story?since=&until=&kind=&limit=   StoryList
//	GET /api/v1/characters/{id}/story/digest?date=YYYY-MM-DD      StoryDigest
//
// Kept as events for a few days (one JSON line each, a file a day), then as a
// day's digest: what a chat's answers were judged and what was visited, kept
// for good. See spec/story-spec.yaml.

// StoryEvent is one thing a character did, or that happened to them.
type StoryEvent struct {
	// At is when, RFC 3339.
	At string `yaml:"at" json:"at"`
	// Character is whose story it is (a character's id).
	Character string `yaml:"character" json:"character"`
	// Kind is what happened: one of the StoryKind* constants.
	Kind string `yaml:"kind" json:"kind"`
	// From and To are where it went from and to, when it moved.
	From *StoryPlace `yaml:"from,omitempty" json:"from,omitempty"`
	To   *StoryPlace `yaml:"to,omitempty" json:"to,omitempty"`
	// Ref is what it is about: a turn's, an act's or a plan's name.
	Ref string `yaml:"ref,omitempty" json:"ref,omitempty"`
	// Device is the device it came from, when known.
	Device string `yaml:"device,omitempty" json:"device,omitempty"`
	// Detail is what else the kind says (a question, a view, an outcome).
	Detail map[string]string `yaml:"detail,omitempty" json:"detail,omitempty"`
}

// StoryPlace is where a character was: a want's or a thing's tile, or a cell
// of the canvas, or a view.
type StoryPlace struct {
	// Kind is want, thing, cell or view.
	Kind string `yaml:"kind" json:"kind"`
	ID   string `yaml:"id,omitempty" json:"id,omitempty"`
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// Type is a want's type (or a thing's catalog).
	Type string `yaml:"type,omitempty" json:"type,omitempty"`
	X    *int   `yaml:"x,omitempty" json:"x,omitempty"`
	Y    *int   `yaml:"y,omitempty" json:"y,omitempty"`
}

// What happened in a story.
const (
	// StoryKindCanvasMove: moved on the canvas and stopped (not every step).
	StoryKindCanvasMove = "canvas.move"
	// StoryKindWantArrive / Leave: stepped onto / off a want's or thing's tile.
	StoryKindWantArrive = "want.arrive"
	StoryKindWantLeave  = "want.leave"
	// StoryKindWantOpen / Close: a want's details opened / closed.
	StoryKindWantOpen  = "want.open"
	StoryKindWantClose = "want.close"
	// StoryKindWantExpand: a want's card opened large.
	StoryKindWantExpand = "want.expand"
	// StoryKindViewChange: a view changed (a tab, the canvas's mode).
	StoryKindViewChange = "view.change"
	// StoryKindSay: said something (the character's speech).
	StoryKindSay = "say"
	// StoryKindChatAsk: asked the robot (Ref, the turn).
	StoryKindChatAsk = "chat.ask"
	// StoryKindChatAnswer: was answered (Ref, the turn; Detail cards, the
	// wants shown: "kind:id" comma-separated).
	StoryKindChatAnswer = "chat.answer"
	// StoryKindActDone: a plan was carried out (Ref, the act).
	StoryKindActDone = "act.done"
	// StoryKindTurnOutcome: an answer judged by what was done next (Ref, the
	// turn; Detail outcome, one of the TurnOutcome* constants).
	StoryKindTurnOutcome = "turn.outcome"
)

// How an answer went, judged by what the person did after it.
const (
	// TurnOutcomeSuccess: went to a want the answer showed.
	TurnOutcomeSuccess = "success"
	// TurnOutcomeCorrected: went to another want of the same type instead —
	// the one they went to was the answer.
	TurnOutcomeCorrected = "corrected"
	// TurnOutcomeReask: asked again soon after.
	TurnOutcomeReask = "reask"
	// TurnOutcomeIgnored: nothing followed from it.
	TurnOutcomeIgnored = "ignored"
)

// StoryList is GET /characters/{id}/story: events, oldest first.
type StoryList struct {
	Events []StoryEvent `yaml:"events" json:"events"`
}

// StoryDigest is a day of a character's story, kept for good after its events
// are let go: each answer and how it went, and what was visited.
type StoryDigest struct {
	Character string `yaml:"character" json:"character"`
	// Date is the day, YYYY-MM-DD in the server's time.
	Date string `yaml:"date" json:"date"`
	// Turns are the day's questions, each with how it went.
	Turns []StoryTurn `yaml:"turns" json:"turns"`
	// Transitions are moves from one want or thing to another, counted.
	Transitions []StoryTransition `yaml:"transitions" json:"transitions"`
	// Visits are the wants and things visited, with how long and how often.
	Visits []StoryVisit `yaml:"visits" json:"visits"`
	// Counts are the day's events by kind, and its outcomes by outcome.
	Counts   map[string]int `yaml:"counts" json:"counts"`
	Outcomes map[string]int `yaml:"outcomes" json:"outcomes"`
	// Summary is the day in a sentence or two, when a model wrote one
	// (SummaryBy says which).
	Summary   string `yaml:"summary,omitempty" json:"summary,omitempty"`
	SummaryBy string `yaml:"summaryBy,omitempty" json:"summaryBy,omitempty"`
}

// StoryTurn is one question of a day, and how its answer went.
type StoryTurn struct {
	Turn     string `yaml:"turn" json:"turn"`
	At       string `yaml:"at" json:"at"`
	Question string `yaml:"question" json:"question"`
	// Plan is what was decided for it, when an act carried one out.
	Plan *PlanSpec `yaml:"plan,omitempty" json:"plan,omitempty"`
	// Shown are the wants the answer showed ("kind:id").
	Shown []string `yaml:"shown,omitempty" json:"shown,omitempty"`
	// Outcome is how it went (TurnOutcome*); CorrectedTo, for corrected, the
	// want gone to instead.
	Outcome     string      `yaml:"outcome,omitempty" json:"outcome,omitempty"`
	CorrectedTo *StoryPlace `yaml:"correctedTo,omitempty" json:"correctedTo,omitempty"`
	// SecondsToAct is how long after the answer the person went somewhere.
	SecondsToAct int `yaml:"secondsToAct,omitempty" json:"secondsToAct,omitempty"`
	// Intent is the question put in a canonical form, when a model wrote one
	// (IntentBy says which): "transit_search: 中野坂上→戸塚, time=now".
	Intent   string `yaml:"intent,omitempty" json:"intent,omitempty"`
	IntentBy string `yaml:"intentBy,omitempty" json:"intentBy,omitempty"`
}

// StoryTransition is a move from one want or thing to another, counted.
type StoryTransition struct {
	From  StoryPlace `yaml:"from" json:"from"`
	To    StoryPlace `yaml:"to" json:"to"`
	Count int        `yaml:"count" json:"count"`
}

// StoryVisit is a want or thing visited in a day.
type StoryVisit struct {
	Place   StoryPlace `yaml:"place" json:"place"`
	Arrived int        `yaml:"arrived" json:"arrived"`
	Opened  int        `yaml:"opened" json:"opened"`
	// Seconds is how long the character stood on it, in all.
	Seconds int `yaml:"seconds" json:"seconds"`
}
