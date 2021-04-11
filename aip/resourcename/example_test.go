package resourcename

import "fmt"

// This example illustrates how to take the name of a parent resource,
// add a resource ID for a child resource, and render the child
// resource name.
func Example() {
	p := MustCompile("blurbs/{blurb}/gobs/{gob}")

	parent := "blurbs/123"
	v, err := p.Match(parent)
	if err != nil {
		fmt.Println(err)
	}

	v["gob"] = "456"
	child, err := p.Render(v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(child)
	// Output:
	// blurbs/123/gobs/456
}

func ExampleCompile_noVariables() {
	p, err := Compile("config")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("p = %q\n", p)
	// Output:
	// p = "config"
}

func ExampleCompile_oneVariable() {
	p, err := Compile("blurbs/{blurb}")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("p = %q\n", p)
	// Output:
	// p = "blurbs/{blurb}"
}

func ExamplePattern_Match_full() {
	p := MustCompile("blurbs/{blurb}")
	v, err := p.Match("blurbs/123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("blurb = %q\n", v["blurb"])
	// Output:
	// blurb = "123"
}

func ExamplePattern_Match_prefix() {
	p := MustCompile("blurbs/{blurb}/gobs/{gob}")
	v, err := p.Match("blurbs/123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("blurb = %q\n", v["blurb"])
	_, ok := v["gob"]
	fmt.Printf("gob exists = %v\n", ok)
	// Output:
	// blurb = "123"
	// gob exists = false
}

func ExamplePattern_Matches() {
	p := MustCompile("blurbs/{blurb}/gobs/{gob}")
	for _, name := range []string{
		"blurbs",
		"blurbs/123",
		"blurbs/{blurb}",
		"blurbs/123/gobs/456",
		"blurbs/123/gobs/{gob}",
	} {
		fmt.Printf("p.Matches(%q) = %v\n", name, p.Matches(name))
	}
	// Output:
	// p.Matches("blurbs") = true
	// p.Matches("blurbs/123") = true
	// p.Matches("blurbs/{blurb}") = false
	// p.Matches("blurbs/123/gobs/456") = true
	// p.Matches("blurbs/123/gobs/{gob}") = false
}

func ExamplePattern_Render() {
	p := MustCompile("blurbs/{blurb}/gobs/{gob}")

	// Set values for all variables.
	v := Values{}
	v["blurb"] = "123"
	v["gob"] = "456"
	// Any extra values are ignored.
	v["rubbish"] = "foobar"

	name, err := p.Render(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)
	// Output:
	// blurbs/123/gobs/456
}

func ExamplePattern_String() {
	p := MustCompile("blurbs/{blurb}/gobs/{gob}")
	fmt.Println(p.String())
	// Output:
	// blurbs/{blurb}/gobs/{gob}
}
