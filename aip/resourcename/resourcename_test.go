package resourcename

import (
	"testing"

	"github.com/Saser/pdp/testing/errtest"
	"github.com/google/go-cmp/cmp"
)

func TestCompile(t *testing.T) {
	patterns := []string{
		"singleton",
		"singleton/blurb",
		"blurbs/{blurb}",
		"blurbs/{blurb}/singleton",
		"blurbs/{blurb}/gobs/{gob}",
		"blurbs/{blurb}/{otherblurb}",
	}
	for _, pattern := range patterns {
		p, err := Compile(pattern)
		if err != nil {
			t.Errorf("Compile(%q) err = %v; want nil", pattern, err)
		}
		if p == nil {
			t.Errorf("Compile(%q) p = nil; want non-nil", pattern)
		}
	}
	t.Run("MustCompile", func(t *testing.T) {
		for _, pattern := range patterns {
			func() {
				defer func() {
					if v := recover(); v != nil {
						t.Errorf("MustCompile(%q) panicked; recover() = %v", pattern, v)
					}
				}()
				p := MustCompile(pattern)
				if p == nil {
					t.Errorf("MustCompile(%q) = nil; want non-nil", pattern)
				}
			}()
		}
	})
}

func TestCompile_Errors(t *testing.T) {
	patterns := []string{
		"",
		"/",
		"foo/",
		"/foo",
		"foo//bar",
		"foo/{bar",
		"foo/bar}",
		"foo/bad{bar}",
		"foos/{foo}/bars/{foo}",
	}
	for _, pattern := range patterns {
		_, err := Compile(pattern)
		if err == nil {
			t.Errorf("Compile(%q) err = nil; want non-nil", pattern)
		}
	}

	t.Run("MustCompile", func(t *testing.T) {
		for _, pattern := range patterns {
			func() {
				defer func() {
					if v := recover(); v == nil {
						t.Errorf("MustCompile(%q) didn't panic; recover() = nil", pattern)
					}
				}()
				MustCompile(pattern)
			}()
		}
	})
}

func TestPattern_Match(t *testing.T) {
	for _, tt := range []struct {
		testName string
		p        *Pattern
		values   map[string]Values
	}{
		{
			testName: "NestedResources",
			p:        MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			values: map[string]Values{
				"blurbs": {},
				"blurbs/123": {
					"blurb": "123",
				},
				"blurbs/123/gobs": {
					"blurb": "123",
				},
				"blurbs/123/gobs/456": {
					"blurb": "123",
					"gob":   "456",
				},
				"blurbs/123/gobs/456/789": {
					"blurb":    "123",
					"gob":      "456",
					"othergob": "789",
				},
			},
		},
		{
			testName: "VariousValues",
			p:        MustCompile("{v}"),
			values: map[string]Values{
				"-":                                    {"v": "-"},
				"eb62bf24-dac0-4e4e-867b-a2219aebd0e8": {"v": "eb62bf24-dac0-4e4e-867b-a2219aebd0e8"},
				"camelCase":                            {"v": "camelCase"},
				"les-miserables":                       {"v": "les-miserables"},
			},
		},
		{
			testName: "OnlySingletons",
			p:        MustCompile("some/nested/singletons"),
			values: map[string]Values{
				"some":                   {},
				"some/nested":            {},
				"some/nested/singletons": {},
			},
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			for name, want := range tt.values {
				got, err := tt.p.Match(name)
				if err != nil {
					t.Errorf("p.Match(%q) err = %v; want nil", name, err)
				}
				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("p.Match(%q): unexpected Values (-want +got)\n%s", name, diff)
				}
			}
			if t.Failed() {
				t.Logf("p = %q", tt.p)
			}
		})
	}
}

func TestPattern_Match_Errors(t *testing.T) {
	for _, tt := range []struct {
		testName string
		p        *Pattern
		names    []string
	}{
		{
			testName: "NestedResources",
			p:        MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			names: []string{
				"",
				"/",
				"blurbs/",
				"/gobs",
				"blurbs/{blurb}",
				"blurbs/123/gobs/{gob}",
				"blurbs/{blurb}/gobs/456",
				"blurbs/{blurb}/gobs/{gob}/789",
				"blurbs/123/gobs/456/{othergob}",
				"blurbs//gobs/456/789",
				"blurbs/123/gobs//789",
				"blurbs/123/gobs/456/",
			},
		},
		{
			testName: "OnlySingletons",
			p:        MustCompile("some/nested/singletons"),
			names: []string{
				"",
				"/",
				"some/",
				"some/nested/",
				"/nested/singletons",
				"some/nested/wrong",
			},
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			for _, name := range tt.names {
				_, err := tt.p.Match(name)
				if err == nil {
					t.Errorf("p.Match(%q) err = nil; want non-nil", name)
				}
			}
			if t.Failed() {
				t.Logf("p = %q", tt.p)
			}
		})
	}
}

func TestPattern_Matches(t *testing.T) {
	for _, tt := range []struct {
		testName string
		p        *Pattern
		cases    map[string]bool
	}{
		{
			testName: "NestedResources",
			p:        MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			cases: map[string]bool{
				"":                               false,
				"blurbs":                         true,
				"blurbs/":                        false,
				"blurbs/123":                     true,
				"blurbs/{blurb}":                 false,
				"blurbs/123/gobs/456/789":        true,
				"blurbs/{blurb}/gobs/456/789":    false,
				"blurbs/123/gobs/{gob}/789":      false,
				"blurbs/123/gobs/456/{othergob}": false,
			},
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			for name, want := range tt.cases {
				got := tt.p.Matches(name)
				if got != want {
					t.Errorf("p.Matches(%q) = %v; want %v", name, got, want)
				}
			}
			if t.Failed() {
				t.Logf("p = %q", tt.p)
			}
		})
	}
}

func TestPattern_Render(t *testing.T) {
	for _, tt := range []struct {
		name string
		p    *Pattern
		v    Values
		want string
	}{
		{
			name: "Singleton_EmptyValues",
			p:    MustCompile("singleton"),
			v:    Values{},
			want: "singleton",
		},
		{
			name: "Singleton_RubbishValues",
			p:    MustCompile("singleton"),
			v: Values{
				"rubbish": "foo",
				"bogus":   "bar",
			},
			want: "singleton",
		},
		{
			name: "NestedSingleton_EmptyValues",
			p:    MustCompile("some/nested/singleton"),
			v:    Values{},
			want: "some/nested/singleton",
		},
		{
			name: "NestedSingleton_RubbishValues",
			p:    MustCompile("some/nested/singleton"),
			v: Values{
				"rubbish": "foo",
				"bogus":   "bar",
			},
			want: "some/nested/singleton",
		},
		{
			name: "TopLevelResource",
			p:    MustCompile("blurbs/{blurb}"),
			v: Values{
				"blurb": "123",
			},
			want: "blurbs/123",
		},
		{
			name: "SubResource",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}"),
			v: Values{
				"blurb": "123",
				"gob":   "456",
			},
			want: "blurbs/123/gobs/456",
		},
		{
			name: "SubResource_ExtraValues",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}"),
			v: Values{
				"blurb":   "123",
				"gob":     "456",
				"rubbish": "foo",
				"bogus":   "bar",
			},
			want: "blurbs/123/gobs/456",
		},
		{
			name: "DeeperSubResource",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v: Values{
				"blurb":    "123",
				"gob":      "456",
				"othergob": "789",
			},
			want: "blurbs/123/gobs/456/789",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Render(tt.v)
			if err != nil {
				t.Errorf("p.Render(%q) err = %v; want nil", tt.v, err)
			}
			if got != tt.want {
				t.Errorf("p.Render(%q) name = %q; want %q", tt.v, got, tt.want)
			}
			if t.Failed() {
				t.Logf("p = %q", tt.p)
			}
		})
	}
}

func TestPattern_Render_Errors(t *testing.T) {
	for _, tt := range []struct {
		name string
		p    *Pattern
		v    Values
		tf   errtest.TestFunc
	}{
		{
			name: "EmptyValues",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v:    Values{},
			tf: errtest.All(
				errtest.ErrorContains("missing value"),
				errtest.ErrorContains(`"blurb"`),
				errtest.ErrorContains(`"gob"`),
				errtest.ErrorContains(`"othergob"`),
			),
		},
		{
			name: "SomeMissingValues",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v: Values{
				"blurb": "123",
			},
			tf: errtest.All(
				errtest.ErrorContains("missing value"),
				errtest.ErrorContains(`"gob"`),
				errtest.ErrorContains(`"othergob"`),
			),
		},
		{
			name: "SomeMissingValues_WithExtraValues",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v: Values{
				"blurb":   "123",
				"rubbish": "foo",
				"bogus":   "bar",
			},
			tf: errtest.All(
				errtest.ErrorContains("missing value"),
				errtest.ErrorContains(`"gob"`),
				errtest.ErrorContains(`"othergob"`),
			),
		},
		{
			name: "VariableSetButEmpty",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v: Values{
				"blurb":    "",
				"gob":      "456",
				"othergob": "789",
			},
			tf: errtest.All(
				errtest.ErrorContains("empty value"),
				errtest.ErrorContains(`"blurb"`),
			),
		},
		{
			name: "VariableSetButInvalid",
			p:    MustCompile("blurbs/{blurb}/gobs/{gob}/{othergob}"),
			v: Values{
				"blurb":    "an_invalid_value",
				"gob":      "456",
				"othergob": "789",
			},
			tf: errtest.All(
				errtest.ErrorContains("invalid_value"),
				errtest.ErrorContains(`"blurb"`),
				errtest.ErrorContains(`"an_invalid_value"`),
			),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.p.Render(tt.v)
			if err == nil {
				t.Errorf("p.Render(%q) err = nil; want non-nil", tt.v)
			}
			tt.tf(t, err)
			if t.Failed() {
				t.Logf("p = %q", tt.p)
			}
		})
	}
}
