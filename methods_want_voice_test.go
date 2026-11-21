package want_spec

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestWantSpecVoice(t *testing.T) {
	var s WantSpec
	if err := yaml.Unmarshal([]byte("params: {}\nvoice:\n  murmur: [soon]\n  notify: []\n"), &s); err != nil {
		t.Fatal(err)
	}
	if len(s.UnknownFields) != 0 || s.Voice == nil || s.Voice.Murmur == nil || (*s.Voice.Murmur)[0] != "soon" ||
		s.Voice.Notify == nil || len(*s.Voice.Notify) != 0 {
		t.Fatalf("yaml voice = %+v (unknown %v)", s.Voice, s.UnknownFields)
	}
	// An empty list (none) survives the round trip; a list left out stays out.
	b, _ := json.Marshal(s)
	var back WantSpec
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Voice == nil || back.Voice.Notify == nil || len(*back.Voice.Notify) != 0 || (*back.Voice.Murmur)[0] != "soon" {
		t.Fatalf("json round trip: %s → %+v", b, back.Voice)
	}
	var none WantSpec
	_ = json.Unmarshal([]byte(`{"params":{},"voice":{"murmur":["a"]}}`), &none)
	if none.Voice.Notify != nil {
		t.Fatal("a list left out came back set")
	}
}
