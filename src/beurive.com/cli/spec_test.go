package cli

// Run the command below to activate the printing of the messages.
// go test -test.v

import (
	"testing"
	"fmt"
	"reflect"
	"strings"
)

// -----------------------------------------------------------------
// Test the case where the specification is OK.
// -----------------------------------------------------------------

func TestSpecInitOk(t *testing.T)  {
	var cloVerbose bool
	var cloInput string
	var cloOutput string
	var cloPaths []string
	var cloLevel int

	os := []string{"v", "i", "o", "l"}
	ol := []string{"verbose", "input", "output", "path"}

	var spec = Spec{
		Option{Short: "i", Long: "input",   Holder: &cloInput},
		Option{Short: "o", Long: "output",  Holder: &cloOutput},
		Option{Short: "",  Long: "path",    Holder: &cloPaths},
		Option{Short: "l", Long: "",        Holder: &cloLevel},
		Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
	}

	index, err := spec.init()

	if nil != err {
		t.Error("The CLI specification should be OK.")
	}

	// Check the number of indexed short options.

	if 4 != len(index.Short) {
		t.Error(fmt.Sprintf(`The number of short options specifiers should be 4. Go %d specifiers.`, len(index.Short)));
	}

	// Check that all short option specifiers are indexed.

	for _, s := range os {
		if _, exits := index.Short[s]; ! exits {
			t.Error(fmt.Sprintf(`Option specifier "%c" not found in the index. It should be found.`, s))
		}
	}

	// Check the types of all short option specifiers holders.

	if _, ok := index.Short["v"].Holder.(*bool); ! ok {
		t.Error(fmt.Sprintf(`The type of the holder associated to the short option "-v" should be *bool. Got %s`, reflect.TypeOf(index.Short["v"].Holder)))
	}

	if _, ok := index.Short["i"].Holder.(*string); ! ok {
		t.Error(`The type of the holder associated to the short option "-i" should be *string.`)
	}

	if _, ok := index.Short["o"].Holder.(*string); ! ok {
		t.Error(`The type of the holder associated to the short option "-o" should be *string.`)
	}

	if _, ok := index.Short["l"].Holder.(*int); ! ok {
		t.Error(`The type of the holder associated to the short option "-l" should be *int.`)
	}

	// Check the number of indexed long options.

	if 4 != len(index.Long) {
		t.Error(fmt.Sprintf(`The number of long options specifiers should be 4. Go %d specifiers.`, len(index.Long)));
	}

	// Check that all long option specifiers are indexed.

	for _, s := range ol {
		if _, exits := index.Long[s]; ! exits {
			t.Error(fmt.Sprintf(`Option specifier "%c" not found in the index. It should be found.`, s))
		}
	}

	// Check the types of all long option specifiers holders.

	if _, ok := index.Long["verbose"].Holder.(*bool); ! ok {
		t.Error(`The type of the holder associated to the long option "--verbose" should be *bool.`)
	}

	if _, ok := index.Long["input"].Holder.(*string); ! ok {
		t.Error(`The type of the holder associated to the long option "--input" should be *string.`)
	}

	if _, ok := index.Long["output"].Holder.(*string); ! ok {
		t.Error(`The type of the holder associated to the long option "--ouput" should be *string.`)
	}

	if _, ok := index.Long["path"].Holder.(*[]string); ! ok {
		t.Error(`The type of the holder associated to the long option "--path" should be []*string.`)
	}
}

// -----------------------------------------------------------------
// Test the detection of a value holder that is used by two distinct
// options.
// -----------------------------------------------------------------

func TestEM_InvalidCmdLineSpecReuseOfValueHolder(t *testing.T)  {
	var cloVerbose bool
	var cloPath string
	var cloPaths []string
	var cloLevel int

	var spec = Spec{
		Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
		Option{Short: "i", Long: "input",   Holder: &cloPath},
		Option{Short: "o", Long: "output",  Holder: &cloPath},
		Option{Short: "",  Long: "path",    Holder: &cloPaths},
		Option{Short: "l", Long: "",        Holder: &cloLevel},
	}

	if _, err := spec.init(); nil == err {
		t.Error(`The specification should not be valid! One value holder is shared between 2 options.`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecReuseOfValueHolder, 2)
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of shared use of name for long options.
// -----------------------------------------------------------------

func TestEM_InvalidCmdLineSpecDuplicatedLongNamedOption(t *testing.T)  {
	var cloVerbose bool
	var cloPath string
	var cloPaths []string
	var cloLevel int

	var spec = Spec{
		Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
		Option{Short: "i", Long: "input",   Holder: &cloPath},
		Option{Short: "",  Long: "input",   Holder: &cloPaths},
		Option{Short: "l", Long: "",        Holder: &cloLevel},
	}

	if _, err := spec.init(); nil == err {
		t.Error(`The specification should not be valid! One value holder is shared between 2 options.`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecDuplicatedLongNamedOption, 2, "input")
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

// -----------------------------------------------------------------
// Test the detection of shared use of name for short options.
// -----------------------------------------------------------------

func TestEM_InvalidCmdLineSpecDuplicatedShortNamedOption(t *testing.T)  {
	var cloVerbose bool
	var cloPath string
	var cloPaths []string
	var cloLevel int

	var spec = Spec{
		Option{Short: "v", Long: "verbose", Holder: &cloVerbose},
		Option{Short: "i", Long: "input",   Holder: &cloPath},
		Option{Short: "i", Long: "inputs",  Holder: &cloPaths},
		Option{Short: "l", Long: "",        Holder: &cloLevel},
	}

	if _, err := spec.init(); nil == err {
		t.Error(`The specification should not be valid! One value holder is shared between 2 options.`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecDuplicatedShortNamedOption, 2, "i")
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

// -----------------------------------------------------------------
// Test the error while initializing an option.
// -----------------------------------------------------------------

func TestEM_InvalidCmdLineSpecInvalidOptionDefinition(t *testing.T)  {
	var cloTypeOk bool
	var cloTypeKo typeOption
	var spec Spec

	// An option without names.
	spec = Spec{
		Option{Short: "", Long: "", Holder: &cloTypeOk},
	}
	if _, err := spec.init(); nil == err {
		t.Error(`Unexpected error: the test should fail!`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, 0, errorInvalidOptionSpecificationNoName)
		if 0 != strings.Compare(m, err.Error()) {
			t.Errorf(`Expected: %s\nGot: %s`, m, err.Error())
		}
	}

	// An Option with a short name that contains more than one character.
	spec = Spec{
		Option{Short: "vv", Long: "", Holder: &cloTypeOk},
	}
	if _, err := spec.init(); nil == err {
		t.Error(`Unexpected error: the test should fail!`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, 0, fmt.Sprintf(errorShortNameTooLong, "vv"))
		if 0 != strings.Compare(m, err.Error()) {
			t.Errorf("Expected: %s\nGot: %s", m, err.Error())
		}
	}

	// An Option with a short name that contains an unexpected character.
	spec = Spec{
		Option{Short: ".", Long: "", Holder: &cloTypeOk},
	}
	if _, err := spec.init(); nil == err {
		t.Error(`Unexpected error: the test should fail!`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, 0, fmt.Sprintf(errorShortNameUnexpectedCharacter, "."))
		if 0 != strings.Compare(m, err.Error()) {
			t.Errorf("Expected: %s\nGot: %s", m, err.Error())
		}
	}

	// An option with a long name that contains an unexpected character.
	// An Option with a short name that contains an unexpected character.
	spec = Spec{
		Option{Short: "", Long: "...", Holder: &cloTypeOk},
	}
	if _, err := spec.init(); nil == err {
		t.Error(`Unexpected error: the test should fail!`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, 0, fmt.Sprintf(errorLongNameUnexpectedCharacter, "..."))
		if 0 != strings.Compare(m, err.Error()) {
			t.Errorf("Expected: %s\nGot: %s", m, err.Error())
		}
	}

	// An option which value holder is of an unexpected type.
	spec = Spec{
		Option{Short: "v", Long: "", Holder: &cloTypeKo},
	}
	if _, err := spec.init(); nil == err {
		t.Error(`Unexpected error: the test should fail!`)
	} else {
		m := fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, 0, errorInvalidOptionSpecificationUnexpectedHolderType)
		if 0 != strings.Compare(m, err.Error()) {
			t.Errorf("Expected: %s\nGot: %s", m, err.Error())
		}
	}
}

