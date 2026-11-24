package want_spec

import (
	"strings"
	"testing"
)

func TestContextKinds(t *testing.T) {
	seen := map[string]bool{}
	for _, k := range ContextKinds {
		if seen[k.Name] {
			t.Errorf("%s twice", k.Name)
		}
		seen[k.Name] = true
		if !strings.HasPrefix(k.Name, k.Category+".") || k.Description == "" {
			t.Errorf("%+v", k)
		}
	}
	if _, ok := ContextKindByName("place.destination"); !ok {
		t.Error("place.destination not found")
	}
}
