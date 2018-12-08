package cli

import (
	"testing"
	"strings"
	"fmt"
)

// -----------------------------------------------------------------
// Test the function that detects an option specifier, based on the
// format of the string.
// -----------------------------------------------------------------

func TestIsOptionOk(t *testing.T)  {
	params := []string{"-o", "-vh", "--option"}
	for _, p := range params {
		if ! isOption(p) {
			t.Errorf(`The specifier "%s" should be a valid option specifier!`, p)
		}
	}
}

func TestIsOptionKo(t *testing.T)  {
	params := []string{"o", "vh"}
	for _, p := range params {
		if isOption(p) {
			t.Errorf(`The specifier "%s" should NOT be a valid option specifier!`, p)
		}
	}
}

// -----------------------------------------------------------------
// Test the function that tests if an option specifier represents a
// short option, or a batch of short options.
// -----------------------------------------------------------------

func TestIsOptionShortOk(t *testing.T)  {
	params := map[string][]string{
		"-o":  []string{"o"},
		"o":   []string{"o"},
		"-vh": []string{"v", "h"},
		"vh":  []string{"v", "h"},
	}

	for param, expected := range params {
		if ok, names := isOptionShort(param); ok {
			if len(expected) != len(names) {
				t.Errorf(`Unexpected list of names! Got %d names. Expected %d.`, len(names), len(expected))
			}
			for i, v := range expected {
				if 0 != strings.Compare(v, names[i]) {
					t.Errorf(`Unexpected list of names! "%s" != "%s"`, v, names[i])
				}
			}
		} else {
			t.Errorf(`The string "%s" should represent a short option, or a batch of short options!`, param)
		}
	}
}

func TestIsOptionShortKo(t *testing.T)  {
	params := []string{"--option", "-o-"}

	for _, p := range params {
		// Discard the name of the option
		if ok, _ := isOptionShort(p); ok {
			t.Errorf(`The specifier "%s" should not be valid`, p)
		}
	}
}

// -----------------------------------------------------------------
// Test the function that tests if an option specifier represents a
// long option.
// -----------------------------------------------------------------

func TestIsOptionLongLongOk(t *testing.T)  {
	params := map[string]string{
		"--o": "o",
		"o": "o",
		"o-": "o-",
		"--option": "option",
		"--option-": "option-",
		"--option.": "option.",
		"option": "option",
		"option-": "option-",
		"option.": "option.",
	}

	for specifier, expected := range params {
		if ok, v := isOptionLong(specifier); ok {
			if 0 != strings.Compare(v, expected) {
				t.Errorf(`"%s" != "%s"`, v, expected)
			}
		} else {
			t.Errorf(`The specificier "%s" should specify for a long option.`, specifier)
		}
	}
}

func TestIsOptionLongLongKo(t *testing.T)  {
	params := []string{"-o", "-option", "---option", "--.o", ".o", "--.option"}

	for _, specifier := range params {
		// Discard the name of the option
		if ok, _ := isOptionLong(specifier); ok {
			t.Errorf(`The specificier "%s" should NOT specify for a long option.`, specifier)
		}
	}
}

// -----------------------------------------------------------------
// Test the function that expands a command line.
// -----------------------------------------------------------------

func TestParseOk(t *testing.T)  {

	var cloVerbose       bool
	var cloInput         string
	var cloPath          []string
	var cloP             bool
	var cloHumanReadable bool
	var cloId            int
	var numInt8          int8
	var numUInt8         uint8
	var numsInt          []int
	var numsUInt         []uint

	type typeExpected struct {
		tokens  []string
		options []Option
		args    []string
	}

	type setType struct {
		spec     Spec
		input    []string
		expected typeExpected
	}

	testSet := []setType{
		{
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "p", Long: "path",  Holder: &cloPath},
			},
			input: []string{ "-v", "--input", "/tmp/file.txt", "-p", "/tmp/a", "-p", "/tmp/b", "--path", "/tmp/c" },
			expected: typeExpected{
				tokens: []string{ "-v", "--input", "/tmp/file.txt", "-p", "/tmp/a", "-p", "/tmp/b", "--path", "/tmp/c" },
				options: []Option{
					Option{Short: "",  Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true },
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true },
					Option{Short: "p", Long: "path",  Holder: (func() *[]string { v := []string{"/tmp/a", "/tmp/b", "/tmp/c"}; return &v })(), set: true},
				},
				args: []string{},
			},
		},
		{
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "p", Long: "path",  Holder: &cloPath},
			},
			input: []string{ "-v", "--input", "/tmp/file.txt", "-p", "/tmp/a", "-p", "/tmp/b", "--path", "/tmp/c", "0", "1" },
			expected: typeExpected{
				tokens: []string{ "-v", "--input", "/tmp/file.txt", "-p", "/tmp/a", "-p", "/tmp/b", "--path", "/tmp/c", "0", "1" },
				options: []Option{
					Option{Short: "",  Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true },
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true },
					Option{Short: "p", Long: "path",  Holder: (func() *[]string { v := []string{"/tmp/a", "/tmp/b", "/tmp/c"}; return &v })(), set: true},
				},
				args: []string{"0", "1"},
			},
		},
		{
			spec: Spec{
				Option{Short: "i", Long: "input", Holder: &cloInput},
				Option{Short: "p", Long: "",      Holder: &cloP},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "",  Long: "pv",    Holder: &cloPath},
			},
			// This is a trap: "-v", "-p" and "-pv"
			input: []string{ "-v", "--input", "/tmp/file.txt", "-p", "--pv", "/tmp/path" },
			expected: typeExpected{
				tokens: []string{ "-v", "--input", "/tmp/file.txt", "-p", "--pv", "/tmp/path" },
				options: []Option{
					Option{Short: "i", Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "p", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "",  Long: "pv",    Holder: (func() *[]string { v := []string{"/tmp/path"}; return &v })(), set: true},
				},
				args: []string{},
			},
		},
		{
			spec: Spec{
				Option{Short: "i", Long: "input", Holder: &cloInput},
				Option{Short: "p", Long: "",      Holder: &cloP},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "",  Long: "pv",    Holder: &cloPath},
			},
			// This is a trap: "-pv" and "--pv".
			input: []string{ "--input", "/tmp/file.txt", "-pv", "--pv", "/tmp/path" },
			expected: typeExpected{
				tokens: []string{ "--input", "/tmp/file.txt", "-p", "-v", "--pv", "/tmp/path" },
				options: []Option{
					Option{Short: "i", Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "p", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "",  Long: "pv",    Holder: (func() *[]string { v := []string{"/tmp/path"}; return &v })(), set: true},
				},
				args: []string{},
			},
		},
		{
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "h", Long: "",      Holder: &cloHumanReadable},
			},
			input: []string{ "-vh", "--input", "/tmp/file.txt", "--" },
			expected: typeExpected{
				tokens: []string{ "-v", "-h", "--input", "/tmp/file.txt", "--" },
				options: []Option{
					Option{Short: "",  Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "h", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
				},
				args: []string{},
			},
		},
		{
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "h", Long: "",      Holder: &cloHumanReadable},
			},
			input: []string{ "-vh", "--input", "/tmp/file.txt", "--", "--xyz"},
			expected: typeExpected{
				tokens: []string{ "-v", "-h", "--input", "/tmp/file.txt", "--", "--xyz" },
				options: []Option{
					Option{Short: "",  Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "h", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
				},
				args: []string{"--xyz"},
			},
		},
		{
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
				Option{Short: "h", Long: "",      Holder: &cloHumanReadable},
				Option{Short: "",  Long: "id--",  Holder: &cloId},
			},
			input: []string{ "-vh", "--input", "/tmp/file.txt", "--id--", "00", "--", "--xyz"},
			expected: typeExpected{
				tokens: []string{ "-v", "-h", "--input", "/tmp/file.txt", "--id--", "00", "--", "--xyz" },
				options: []Option{
					Option{Short: "",  Long: "input", Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "v", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "h", Long: "",      Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "",  Long: "id--",  Holder: (func() *int { v := int(0); return &v })(), set: true},
				},
				args: []string{"--xyz"},
			},
		},
		{
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
				Option{Short: "h", Long: "human",   Holder: &cloHumanReadable},
				Option{Short: "",  Long: "id--",    Holder: &cloId},
			},
			input: []string{ "-vh", "--input", "/tmp/file.txt", "--id--", "00", "--", "--xyz"},
			expected: typeExpected{
				tokens: []string{ "-v", "-h", "--input", "/tmp/file.txt", "--id--", "00", "--", "--xyz" },
				options: []Option{
					Option{Short: "i", Long: "input",   Holder: (func() *string { v := "/tmp/file.txt"; return &v })(), set: true},
					Option{Short: "v", Long: "verbose", Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "h", Long: "human",   Holder: (func() *bool { v := true; return &v })(), set: true},
					Option{Short: "",  Long: "id--",    Holder: (func() *int { v := int(0); return &v })(), set: true},
				},
				args: []string{"--xyz"},
			},

		},
		{
			spec: Spec{
				Option{Short: "i", Long: "int8",     Holder: &numInt8},
				Option{Short: "u", Long: "uint8",    Holder: &numUInt8},
				Option{Short: "",  Long: "ints",     Holder: &numsInt},
				Option{Short: "",  Long: "uints",    Holder: &numsUInt},
			},
			input: []string{ "-i", "127", "-u", "255", "--ints", "10", "--ints", "20", "--uints", "100", "--uints", "200", "--", "toto" },
			expected: typeExpected{
				tokens: []string{ "-i", "127", "-u", "255", "--ints", "10", "--ints", "20", "--uints", "100", "--uints", "200", "--", "toto" },
				options: []Option{
					Option{Short: "i", Long: "int8",    Holder: (func() *int8   { v := int8(127);        return &v })(), set: true},
					Option{Short: "u", Long: "uint8",   Holder: (func() *uint8  { v := uint8(255);       return &v })(), set: true},
					Option{Short: "",  Long: "ints",    Holder: (func() *[]int  { v := []int{210, 20};   return &v })(), set: true},
					Option{Short: "",  Long: "uints",   Holder: (func() *[]uint { v := []uint{100, 200}; return &v })(), set: true},
				},
				args: []string{"toto"},
			},
		},
	}

	for i, set := range testSet {

		// Be aware that the same value holders are used for all tests.
		// Since the values are not reset, it is necessary to reset them first.

		cloVerbose = false
		cloInput = ""
		cloPath = []string{}
		cloP = false
		cloHumanReadable = false
		cloId = 0

		if cli, args, err := Parse(set.input, set.spec); nil != err {
			t.Errorf(`The test number %d should be OK (%s). Got the error: %s`, i, strings.Join(set.input, " "), err.Error())
		} else {
			// Check the returned command line against the expected one.
			if len(cli) != len(set.expected.tokens) {
				t.Errorf(`Unexpected list of CLI tokens. Expected (%s) / Got (%s)`, strings.Join(set.expected.tokens, " "), strings.Join(set.input, " "))
			}
			for j, v := range set.expected.tokens {
				if 0 != strings.Compare(cli[j], v) {
					t.Errorf(`Unexpected list of CLI token #%d. Expected (%s) / Got (%s)`, j, strings.Join(set.expected.tokens, " "), strings.Join(set.input, " "))
				}
			}

			// Check the returned list of arguments against the expected one.
			if len(set.expected.args) != len(args) {
				t.Errorf(`Unexpected list of arguments. Expected (%s) / Got (%s)`, strings.Join(set.expected.args, " "), strings.Join(args, " "))
			}

			for j, v := range set.expected.args {
				if 0 != strings.Compare(args[j], v) {
					t.Errorf(`Unexpected argument #%d. Expected (%s) / Got (%s)`, j, strings.Join(set.expected.tokens, " "), strings.Join(set.input, " "))
				}
			}


			// Check the options against the expected ones.
			for j, expected := range set.expected.options {
				t.Logf("=> Test #%d/%d", i, j)

				if set.spec[j].set != expected.set {
					t.Errorf(`Test #%d failed! Got [set=%v], expected [set=%v]`, j, set.spec[j].set, expected.set)
				}
				if expected.set {
					if vv, ok := set.spec[j].Holder.(*bool); ok {
						// Option is a boolean
						expected, _ := set.expected.options[j].Holder.(*bool)
						if *expected != *vv {
							t.Errorf(`Test #%d/%d failed! Got %v, expected %v`, i, j, *expected, *vv)
						}
					} else if vv, ok := set.spec[j].Holder.(*string); ok {
						// Option is a string.
						expected, _ := set.expected.options[j].Holder.(*string)
						if 0 != strings.Compare(*expected, *vv) {
							t.Errorf(`Test #%d/%d failed! Got "%s", expected "%s"`, i, j, *expected, *vv)
						}
					} else if vv, ok := set.spec[j].Holder.(*int); ok {
						// Option is an integer.
						expected, _ := set.expected.options[j].Holder.(*int)
						if *expected != *vv {
							t.Errorf(`Test #%d/%d failed! Got %d, expected %d`, i, j, *expected, *vv)
						}
					} else if vv, ok := set.spec[j].Holder.(*int8); ok {
						// Option is an integer.
						expected, _ := set.expected.options[j].Holder.(*int8)
						if *expected != *vv {
							t.Errorf(`Test #%d/%d failed! Got %d, expected %d`, i, j, *expected, *vv)
						}
					} else if vv, ok := set.spec[j].Holder.(*uint8); ok {
						// Option is an integer.
						expected, _ := set.expected.options[j].Holder.(*uint8)
						if *expected != *vv {
							t.Errorf(`Test #%d/%d failed! Got %d, expected %d`, i, j, *expected, *vv)
						}
					} else if vv, ok := set.spec[j].Holder.(*[]string); ok {
						// Option is a list of strings.
						if expected, ok := set.expected.options[j].Holder.(*[]string); ok {
							if len(*expected) != len(*vv) {
								for _, v := range *vv {
									t.Logf("%#v", v)
								}
								t.Errorf(`Test #%d/%d failed! Got %d elements, expected %d elements`, i, j, len(*vv), len(*expected) )
							}
						} else {
							t.Errorf(`Test #%d/%d failed! Unexpected error: got %#v, expected %#v`, i, j, set.expected.options[j].Holder, expected)
						}
					} else if vv, ok := set.spec[j].Holder.(*[]int); ok {
						// Option is a list of strings.
						if expected, ok := set.expected.options[j].Holder.(*[]int); ok {
							if len(*expected) != len(*vv) {
								for _, v := range *vv {
									t.Logf("%#v", v)
								}
								t.Errorf(`Test #%d/%d failed! Got %d elements, expected %d elements`, i, j, len(*vv), len(*expected) )
							}
						} else {
							t.Errorf(`Test #%d/%d failed! Unexpected error: got %#v, expected %#v`, i, j, set.expected.options[j].Holder, expected)
						}
					} else if vv, ok := set.spec[j].Holder.(*[]uint); ok {
						// Option is a list of strings.
						if expected, ok := set.expected.options[j].Holder.(*[]uint); ok {
							if len(*expected) != len(*vv) {
								for _, v := range *vv {
									t.Logf("%#v", v)
								}
								t.Errorf(`Test #%d/%d failed! Got %d elements, expected %d elements`, i, j, len(*vv), len(*expected) )
							}
						} else {
							t.Errorf(`Test #%d/%d failed! Unexpected error: got %#v, expected %#v`, i, j, set.expected.options[j].Holder, expected)
						}
					} else {
						t.Errorf(`Test #%d/%d failed! Unexpected error (%#v)`, i, j, set.expected.options[j].Holder)
					}
				}
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of unexpected options: The option is not
// declared within the specification.
// -----------------------------------------------------------------

func TestEM_ParseUnexpectedShortNamedOption(t *testing.T)  {
	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Unexpected use of option "-o"
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
			},
			input: []string{ "-v", "--input", "/tmp/file.txt", "-o" },
			expected: fmt.Sprintf(errorUnexpectedShortNamedOption, "o"),
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of unexpected options: The option is not
// declared within the specification.
// -----------------------------------------------------------------

func TestEM_ParseUnexpectedLongNamedOption(t *testing.T)  {
	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Unexpected use of option "-o"
			spec: Spec{
				Option{Short: "",  Long: "input", Holder: &cloInput},
				Option{Short: "v", Long: "",      Holder: &cloVerbose},
			},
			input: []string{ "-v", "--input", "/tmp/file.txt", "--output" },
			expected: fmt.Sprintf(errorUnexpectedLongNamedOption, "output"),
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of duplicated use of flags.
// -----------------------------------------------------------------

func TestEM_ParseDuplicatedFlagOption(t *testing.T)  {

	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Duplicated use of a flag ("-v" ... "-v").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "--input", "/tmp/file.txt", "-v" },
			expected: fmt.Sprintf(errorDuplicatedFlagOption, "v"),
		},
		{
			// Duplicated use of a flag ("-vv").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-vv", "--input", "/tmp/file.txt" },
			expected: fmt.Sprintf(errorDuplicatedFlagOption, "v"),
		},
		{
			// Duplicated use of a flag ("-v ... --verbose").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "--verbose", "--input", "/tmp/file.txt" },
			expected: fmt.Sprintf(errorDuplicatedFlagOption, "verbose"),
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of invalid compounds.
// -----------------------------------------------------------------

func TestEM_ParseNonFlagOptionWithinCompound(t *testing.T)  {

	var cloVerbose bool
	var cloInput string
	var cloOutput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// The option "i" should not be part of a compound.
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
				Option{Short: "o", Long: "output",  Holder: &cloOutput},
			},
			input: []string{ "-vi", "-o", "/tmp/file.txt" },
			expected: fmt.Sprintf(errorNonFlagOptionWithinCompound, "i"),
		},
		{
			// The option "i" should not be part of a compound.
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
				Option{Short: "o", Long: "output",  Holder: &cloOutput},
			},
			input: []string{ "-vio" },
			expected: fmt.Sprintf(errorNonFlagOptionWithinCompound, "i"),
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of duplicated use of singletons.
// -----------------------------------------------------------------

func TestEM_ParseDuplicatedNonFlagOption(t *testing.T)  {

	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Duplicated use of a flag ("-i" ... "--input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "-i", "/tmp/file.txt", "--input", "value" },
			expected: fmt.Sprintf(errorDuplicatedNonFlagOption, "input"),
		},
		{
			// Duplicated use of a flag ("-i" ... "-input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "-i", "/tmp/toto", "--input", "/tmp/file.txt" },
			expected: fmt.Sprintf(errorDuplicatedNonFlagOption, "input"),
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of invalid options specifiers.
// -----------------------------------------------------------------

func TestEM_ParseInvalidOptionSpecifier(t *testing.T)  {

	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Duplicated use of a flag ("-i" ... "--input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "-i", "/tmp/file.txt", "---input", "toto" },
			expected: fmt.Sprintf(errorInvalidOptionSpecifier, "---input", 3),
		},
		{
			// Duplicated use of a flag ("-i" ... "--input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			input: []string{ "-v", "-i", "/tmp/file.txt", "--i!nput" },
			expected: fmt.Sprintf(errorInvalidOptionSpecifier, "--i!nput", 3),
		},

	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the unexpected detection of the string that marks the end of
// the list of options.
// -----------------------------------------------------------------

func TestEM_UnexpectedEndOfOptionsListSpec(t *testing.T)  {
	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Duplicated use of a flag ("-i" ... "--input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			// Option "-i" requires a value (found the end of the list of options).
			input: []string{ "-v", "-i", "--", "0" },
			expected: errorUnexpectedEndOfOptionsListSpec,
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the unexpected detection of an option specifier. A value was
// expected.
// -----------------------------------------------------------------

func TestEM_ValueExpectedOptionEncountered(t *testing.T)  {
	var cloVerbose bool
	var cloInput string

	type setType struct {
		spec Spec
		input []string
		expected string
	}
	testSet := []setType{
		{
			// Duplicated use of a flag ("-i" ... "--input").
			spec: Spec{
				Option{Short: "i", Long: "input",   Holder: &cloInput},
				Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
			},
			// Option "-i" requires a value (found an option).
			input: []string{ "-i", "-v", "--", "0" },
			expected: errorValueExpectedOptionEncountered,
		},
	}

	for i, set := range testSet {
		if _, _, err := Parse(set.input, set.spec); nil == err {
			t.Errorf(`The test number %d should NOT be OK (%s)`, i, strings.Join(set.input, " "))
		} else {
			if 0 != strings.Compare(err.Error(), set.expected) {
				t.Errorf(`Test #%d failed. Got [%s] / [%s]`, i, err.Error(), set.expected)
			}
		}
	}
}

