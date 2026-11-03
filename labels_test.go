package want_spec

import "testing"

func TestLabelSelector(t *testing.T) {
	labels := map[string]string{"device": "iphone", "source": "fmtool"}
	for sel, want := range map[string]bool{
		"":                             true,
		"device=iphone":                true,
		"device==iphone,source":        true,
		"device=mac":                   false,
		"device!=mac":                  true,
		"device!=iphone":               false,
		"missing":                      false,
		"missing!=x":                   true,
		"device=iphone, source=fmtool": true,
	} {
		s, err := ParseLabelSelector(sel)
		if err != nil {
			t.Fatalf("%q: %v", sel, err)
		}
		if got := s.Matches(labels); got != want {
			t.Errorf("%q matches = %v, want %v", sel, got, want)
		}
	}
	if _, err := ParseLabelSelector("=x"); err == nil {
		t.Error("a term with no key is refused")
	}
	if l, err := ParseLabels([]string{"a=1", "b = 2"}); err != nil || l["a"] != "1" || l["b"] != "2" {
		t.Errorf("ParseLabels: %v %v", l, err)
	}
	if _, err := ParseLabels([]string{"nokey"}); err == nil {
		t.Error("a label with no = is refused")
	}
}
