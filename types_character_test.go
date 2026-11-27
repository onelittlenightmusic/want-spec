package want_spec

import "testing"

func TestCharacterTimingsStandardReplacedByName(t *testing.T) {
	got := CharacterTimings([]TimingDef{
		{Name: "uneasy", OnChange: CharacterStateMood, If: CharacterStateMoodUneasy, Notify: true},
		{Name: "moved", OnChange: CharacterStatePlace},
	})
	if len(got) != 3 || got[0].Name != "errand-different" || got[1].Name != "uneasy" || !got[1].Notify || got[2].Name != "moved" {
		t.Fatalf("%+v", got)
	}
	if n := len(CharacterTimings(nil)); n != len(CharacterStandardTimings()) {
		t.Fatalf("%d", n)
	}
}
