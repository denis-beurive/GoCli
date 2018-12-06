package cli


// This package describes a command line.
//
// Anatomy of a command line:
//
// (1) a program name.
// (2) an optional list of options.
// (3) the optional string "--".
// (4) an optional list of arguments.
//
// - An option modifies the operation of the program.
//   * Options are identified by their names.
//   * An option may require a value.
//   * Options that don't require values may be called flags (or switches).
//     Please note that these options have implicit values.
//     If the option appears within the command line, then its implicit value is the boolean value true.
//     Otherwise, its implicit value is the boolean value false.
//   * Options names are prefaced with a single dash (-) or a double dash (--).
//   * Options names prefaced with a single dash are made of a single character (ex: -v).
//     These options are called "short options".
//     Short options can be grouped together (ex: "-vh" is equivalent to "-v -h").
//     A grouped list of short options is called a "compound".
//   * Options names prefaced with a double dash are made of one or more characters (ex: --input).
//     These options are called "long options".
//   * If an option requires a value, this value is separated from the option's name by the character "=" or by one or
//     more space (ex: "-i /path/to/input", "-i=/path/to/input", "--input /path/to/input" or "--input=/path/to/input").
//   * An option that requires values may appear more than once within the command line. In this case, its value will be
//     represented by an array. Please note that all values must be of the same type (all strings, all integers...).
//   * Options that don't require values (called flags or switches) can be grouped (ex: "-v -f" can be written "-vf").
//
// - An argument is an item of information provided to the program.
//   * Contrary to options, arguments are not identified by names. Arguments are identified by their positions.
//   * Arguments should not start with a dash (-). However, this statement is not a mandatory requirement.
//   * If an argument starts with a dash, then the list of arguments must be explicitly separated from the list of options
//     by a double dash (--).
//
// Examples:
//
//     prg -l /path/to/dir
//     prg -ltr -- -file*.log
//     prg --output /path/to/output *.log
//     prg -vl --output /path/to/output *.log
//     prg -v --output /path/to/output -l FATAL *.log
//     prg --output /path/to/output --verbose *.log
//     prg --output /path/to/output --verbose -- -FAT- *.log

import (
	"regexp"
	"fmt"
	"strings"
	"errors"
)


// Test whether a string starts with the option's prefix "-" or "--".
// If the given string starts with the option's prefix, then the function returns the value true.
// Otherwise, it returns the value false.
// Please note that if the returned value is true, then it does not mean that the string represents a valid option
// specifier.

func isOption(inString string) bool {
	var rx *regexp.Regexp = regexp.MustCompile(`^--?`)
	return rx.MatchString(inString)
}

// Test if a string represents a short option specifier.
// The format of the function's parameter can be:
// - "-o" or "o"
// - "-vh" or "vh" (for "-v -h")
// If the given string represents a short option specifier, then the function returns the status true, followed by the
// names of the options. Otherwise, the function returns the status false, followed by the value nil.
// Please note that this function is the reference that defines the format of a short option.

func isOptionShort(inString string) (bool, []string) {
	var rx *regexp.Regexp = regexp.MustCompile(`(?i)^(-)?([!?a-z0-9]+)$`)
	m := rx.FindAllStringSubmatch(inString, -1)
	if nil == m { return false, nil }
	if 1 == len(m[0][2]) {
		return true, []string{m[0][2]}
	}

	res := make([]string, 0)
	// If necessary, split a compound into its components (ex: "-vh" => "-v -h").
	for _, c := range m[0][2] {
		res = append(res, fmt.Sprintf("%c", c))
	}
	return true, res
}

// Test whether a string represents a long option specifier or not.
// The format of the function's parameter can be "--option" or "option".
// If the given string represents a short option specifier, then the function returns the status true, followed by the
// name of the option. Otherwise, the function returns the status false, followed by the value nil.
// Please note that this function is the reference that defines the format of a long option.

func isOptionLong(inString string) (bool, string) {
	var rx *regexp.Regexp = regexp.MustCompile(`(?i)^(--)?([!?a-z0-9][a-z0-9_\-.]*)$`)
	m := rx.FindAllStringSubmatch(inString, -1)
	if nil == m { return false, "" }
	return true, m[0][2]
}

// Identify a flag as being used.

func recordOption(inOption *Option, inName string) error {
	if (! inOption.requireValue()) || inOption.isSingleton() {
		// This is a flag of a singleton. Thus, it can appear only once within the command line.
		if inOption.isSet() {
			var m string
			if ! inOption.requireValue() {
				m = fmt.Sprintf(errorDuplicatedFlagOption, inName)
			} else {
				m = fmt.Sprintf(errorDuplicatedNonFlagOption, inName)
			}
			return errors.New(m)
		}
		inOption.setIt()
		return nil
	}

	inOption.setIt()
	return nil
}

// Test whether a given string represents the sequence of characters that marks the end of the list of options.

func isEndOfOptionSpecifier(inString string) bool {
	return 0 == strings.Compare("--", inString)
}

// Expand a command line, relatively to a given specification.
// In addition to expanding the command line, the function also performs some checking:
// - Check that options that appear within compounds are flags.
// - Check that all options are specified.
// - Check that all options names are valid.
//
// For example, let's consider the command line below:
// command -vh --input /tmp/file -o /tmp/result 123 E2
// It will be expended into:
// command -v -h --input /tmp/file -o /tmp/result 123 E2
//
// The function returns the following elements:
// - A list of strings that represents the expanded command line (the options and the arguments).
// - A list of strings that represents the arguments.
// - An error message, if an error occurred.

func Parse(inCliParams []string, inSpec Spec) (cli []string, args []string, err error) {

	cliAll := make([]string, 0)
	index, error := inSpec.init()
	if nil != error {
		cli = nil
		args = nil
		err = errors.New(error.Error())
		return
	}

	// The value of "lastOption" is not nil if the value of "nextShouldBeValue" is true.
	nextShouldBeValue := false;
	var lastOption *Option

	for i, param := range inCliParams {
		// Test whether we need to find an option's value.
		if nextShouldBeValue {
			if isEndOfOptionSpecifier(param) {
				cli = nil
				args = nil
				err = errors.New(errorUnexpectedEndOfOptionsListSpec)
				return
			}

			if isOption(param) {
				cli = nil
				args = nil
				err = errors.New(errorValueExpectedOptionEncountered)
				return
			}

			// This is an option's value
			cliAll = append(cliAll, param)
			lastOption.addValue(param)
			nextShouldBeValue = false
			continue
		}

		// The string found may be an option specifier, the string that marks the end of the list of options, or
		// an argument.

		// Test whether the string is the string that marks the end of the list of options, or not.
		if isEndOfOptionSpecifier(param) {
			cli = append(cliAll, inCliParams[i:]...)
			args = inCliParams[i+1:]
			err = nil
			return
		}

		// The string may be an option specifier.
		if isOption(param) {
			if ok, options := isOptionShort(param); ok {
				// The specifier may be a compound.
				// If so:
				// - All options must be defined within the specification.
				// - All options must be flags (or switches): they don't require values.
				// Please note that a flag must appear only once within the entire command line.
				for _, name := range options {
					o := index.getShortByName(name)
					if nil == o {
						cli = nil
						args = nil
						err = errors.New(fmt.Sprintf(errorUnexpectedShortNamedOption, name))
						return
					}
					if o.requireValue() && len(options) > 1 {
						cli = nil
						args = nil
						err = errors.New(fmt.Sprintf(errorNonFlagOptionWithinCompound, name))
						return
					}
					if err = recordOption(o, name); nil != err {
						cli = nil
						args = nil
						return
					}
					o.addValue(true)
					cliAll = append(cliAll, fmt.Sprintf(`-%s`, name))
				}
				if (len(options) > 1) {
					// We found a compound.
					nextShouldBeValue = false
					lastOption = nil
				} else {
					// We found an isolated short option.
					o := index.getShortByName(options[0])
					nextShouldBeValue = o.requireValue()
					if nextShouldBeValue {
						lastOption = o
					} else {
						lastOption = nil
					}
				}

			} else if ok, name := isOptionLong(param); ok {
				o := index.getLongByName(name)
				if nil == o {
					cli = nil
					args = nil
					err = errors.New(fmt.Sprintf(errorUnexpectedLongNamedOption, name))
					return
				}
				if err = recordOption(o, name); nil != err {
					cli = nil
					args = nil
					return
				}
				nextShouldBeValue = o.requireValue()
				if nextShouldBeValue {
					lastOption = o
				} else {
					// This is a flag (that does not require a value)
					lastOption = nil
					o.addValue(true)
				}

				cliAll = append(cliAll, fmt.Sprintf(`--%s`, name))
			} else {
				cli = nil
				args = nil
				err = errors.New(fmt.Sprintf(errorInvalidOptionSpecifier, param, i))
				return
			}

			continue
		}

		// At this point, the string represents an argument.
		cli = append(cliAll, inCliParams[i:]...)
		args = inCliParams[i:]
		err = nil
		return
	}
	cli = cliAll
	args = []string{}
	err = nil
	return
}

