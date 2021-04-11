// Package resourcename contains types and operations for working with resource names, mostly
// according to how they are defined in https://google.aip.dev/122.
package resourcename

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	literalRE  = regexp.MustCompile("^[[:alpha:]]+$")
	variableRE = regexp.MustCompile(`^\{[[:alpha:]]+\}$`)
	valueRE    = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
)

// segment represents a segment in a pattern (not in a name).
type segment struct {
	Pattern  string // The string this segment was parsed from.
	Variable bool   // Whether this segment represents a variable.
}

// parseSegment parses the given string as a single segment.
func parseSegment(s string) (segment, error) {
	if s == "" {
		return segment{}, errors.New("empty segment")
	}
	switch {
	case literalRE.MatchString(s):
		return segment{
			Pattern:  s,
			Variable: false,
		}, nil
	case variableRE.MatchString(s):
		return segment{
			Pattern:  s,
			Variable: true,
		}, nil
	default:
		return segment{}, fmt.Errorf("invalid segment %q", s)
	}
}

// VarName returns the name of the variable, i.e., the "blurb" in "{blurb}". If this segment does
// not represent a variable (s.Variable is false), VarName returns an empty string.
func (s segment) VarName() string {
	if !s.Variable {
		return ""
	}
	return strings.Trim(s.Pattern, "{}")
}

// Pattern represents a pattern for resource names. An example pattern can be
// "blurbs/{blurb}/gobs/{gob}", where "blurbs" and "gobs" are literals that must be matched exactly,
// and "{blurb}" and "{gob}" are variables whose values can be replaced. Variables are referred to
// without their surrounding braces, i.e., the variable given in the pattern as "{blurb}" is
// referred to as "blurb".
type Pattern struct {
	segments []segment
}

// Compile parses the given and returns a Pattern that represents the pattern defined by the
// string. The pattern must consist of a non-empty sequence of segments, where each segment is
// either a literal or a variable.
//
// Literals must be non-empty and can only consist of a-z and A-Z. Variables must be non-empty
// identifiers consisting of a-z and A-Z, with a surrounding set of braces. A pattern of
// "blurbs/{blurb}" contains a literal "blurbs" and a variable "blurb".
func Compile(pattern string) (*Pattern, error) {
	if pattern == "" {
		return nil, errors.New("empty pattern")
	}
	segstrs := strings.Split(pattern, "/")
	var segments []segment
	seen := make(map[string]bool)
	for _, segstr := range segstrs {
		seg, err := parseSegment(segstr)
		if err != nil {
			return nil, fmt.Errorf("invalid segment in pattern %q: %v", pattern, err)
		}
		if seg.Variable {
			if seen[segstr] {
				return nil, fmt.Errorf("pattern %q contains duplicate variable %q", pattern, segstr)
			}
			seen[segstr] = true
		}
		segments = append(segments, seg)
	}
	return &Pattern{
		segments: segments,
	}, nil
}

// MustCompile is like Compile, but panics on any error.
func MustCompile(pattern string) *Pattern {
	p, err := Compile(pattern)
	if err != nil {
		panic(err)
	}
	return p
}

// Match matches the given name against this pattern. The variables in the pattern are given values
// based on the corresponding segment in the name. The name can contains the same segments or a
// prefix sequence of segments. However, a name with more segments than the pattern is invalid.
//
// Variable values must be non-empty and can consist only of a-z, A-Z, 0-9, and hyphens.
func (p *Pattern) Match(name string) (Values, error) {
	v := Values{}
	if name == "" {
		return v, nil
	}
	segstrs := strings.Split(name, "/")
	if got, max := len(segstrs), len(p.segments); got > max {
		return Values{}, fmt.Errorf("match: name %q has %d segments, which is more than pattern %q which has %d segments", name, got, p, max)
	}
	for i, segstr := range segstrs {
		seg := p.segments[i]
		if seg.Variable {
			varName := seg.VarName()
			if !valueRE.MatchString(segstr) {
				return Values{}, fmt.Errorf("match: invalid value %q for variable %q", segstr, varName)
			}
			v[varName] = segstr
		} else {
			if segstr != seg.Pattern {
				return Values{}, fmt.Errorf("match: segment %q didn't match literal %q", segstr, seg.Pattern)
			}
		}
	}
	return v, nil
}

// Matches returns whether Match would return a nil error.
func (p *Pattern) Matches(name string) bool {
	_, err := p.Match(name)
	return err == nil
}

// Render uses the given values to construct a resource name. The given values must contain
// non-empty, valid values for all variables in the pattern.
//
// Variable values must be non-empty and can consist only of a-z, A-Z, 0-9, and hyphens.
func (p *Pattern) Render(v Values) (string, error) {
	needed := make(map[string]bool)
	for _, seg := range p.segments {
		if !seg.Variable {
			continue
		}
		needed[seg.VarName()] = true
	}
	var missing []string
	for varName := range needed {
		if _, ok := v[varName]; ok {
			continue
		}
		missing = append(missing, fmt.Sprintf("%q", varName))
	}
	if len(missing) > 0 {
		vars := strings.Join(missing, ", ")
		return "", fmt.Errorf("render: missing value(s) for variable(s) %s", vars)
	}
	var rendered []string
	for _, seg := range p.segments {
		if !seg.Variable {
			rendered = append(rendered, seg.Pattern)
			continue
		}
		varName := seg.VarName()
		val := v[varName]
		if val == "" {
			return "", fmt.Errorf("render: empty value for variable %q", varName)
		}
		if !valueRE.MatchString(val) {
			return "", fmt.Errorf("render: invalid value %q for variable %q", val, varName)
		}
		rendered = append(rendered, val)
	}
	return strings.Join(rendered, "/"), nil
}

// String returns the string that was used to compile this pattern.
func (p *Pattern) String() string {
	var ss []string
	for _, segment := range p.segments {
		ss = append(ss, segment.Pattern)
	}
	return strings.Join(ss, "/")
}

// Values represents a mapping from variables to their values.
type Values map[string]string
