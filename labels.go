package want_spec

import (
	"fmt"
	"strings"
)

// LabelSelector picks objects by their labels, as kubectl's -l does with
// equality: every key=value must be there (key!=value: must not be), and a
// bare key must be there with any value.
//
//	device=iphone,source=fmtool
//	env!=prod,team
type LabelSelector []LabelRequirement

// LabelRequirement is one term of a LabelSelector.
type LabelRequirement struct {
	Key   string
	Op    string // "=", "!=", or "exists"
	Value string
}

// ParseLabelSelector reads "k=v,k2!=v2,k3"; "" selects everything.
func ParseLabelSelector(s string) (LabelSelector, error) {
	var sel LabelSelector
	for _, term := range strings.Split(s, ",") {
		term = strings.TrimSpace(term)
		if term == "" {
			continue
		}
		var r LabelRequirement
		switch {
		case strings.Contains(term, "!="):
			k, v, _ := strings.Cut(term, "!=")
			r = LabelRequirement{Key: strings.TrimSpace(k), Op: "!=", Value: strings.TrimSpace(v)}
		case strings.Contains(term, "=="):
			k, v, _ := strings.Cut(term, "==")
			r = LabelRequirement{Key: strings.TrimSpace(k), Op: "=", Value: strings.TrimSpace(v)}
		case strings.Contains(term, "="):
			k, v, _ := strings.Cut(term, "=")
			r = LabelRequirement{Key: strings.TrimSpace(k), Op: "=", Value: strings.TrimSpace(v)}
		default:
			r = LabelRequirement{Key: term, Op: "exists"}
		}
		if r.Key == "" {
			return nil, fmt.Errorf("bad label selector term %q", term)
		}
		sel = append(sel, r)
	}
	return sel, nil
}

// Matches says whether labels meet every term.
func (sel LabelSelector) Matches(labels map[string]string) bool {
	for _, r := range sel {
		v, ok := labels[r.Key]
		switch r.Op {
		case "=":
			if !ok || v != r.Value {
				return false
			}
		case "!=":
			if ok && v == r.Value {
				return false
			}
		default:
			if !ok {
				return false
			}
		}
	}
	return true
}

// String is the selector as it is written.
func (sel LabelSelector) String() string {
	parts := make([]string, len(sel))
	for i, r := range sel {
		switch r.Op {
		case "exists":
			parts[i] = r.Key
		default:
			parts[i] = r.Key + r.Op + r.Value
		}
	}
	return strings.Join(parts, ",")
}

// ParseLabels reads "k=v" pairs (a --label flag's values) into labels.
func ParseLabels(pairs []string) (map[string]string, error) {
	if len(pairs) == 0 {
		return nil, nil
	}
	out := make(map[string]string, len(pairs))
	for _, p := range pairs {
		k, v, ok := strings.Cut(p, "=")
		k = strings.TrimSpace(k)
		if !ok || k == "" {
			return nil, fmt.Errorf("label %q is not key=value", p)
		}
		out[k] = strings.TrimSpace(v)
	}
	return out, nil
}
