package cli

type errorMessageType string

const (
	// ----------------------------------------------------------------
	// cli.go
	// ----------------------------------------------------------------

	errorDuplicatedFlagOption = `Duplicated use of flag option "%s".`
	errorDuplicatedNonFlagOption = `Duplicated use of non flag option "%s".`
	errorUnexpectedShortNamedOption = `Unexpected option which short name is "%s".`
	errorUnexpectedLongNamedOption = `Unexpected option which long name is "%s".`
	errorNonFlagOptionWithinCompound = `The option "%s" requires a value. It cannot appear within a compound.`
	errorInvalidOptionSpecifier = `Invalid option specifier "%s" at position %d.`
	errorUnexpectedEndOfOptionsListSpec = `Unexpected end of list of options mark ("--") encountered. A value was expected.`
	errorValueExpectedOptionEncountered = `Unexpected option specifier found. A value was expected.`

	// ----------------------------------------------------------------
	// option.go
	// ----------------------------------------------------------------

	errorInvalidOptionSpecificationNoName = `Invalid option specifer: no name (whether short or long) is specified for this option.`
	errorInvalidOptionSpecificationUnexpectedHolderType = `Invalid option definition: wrong type detected for the value holder.`
	errorShortNameTooLong = `Invalid short name "%s". A short name should contain only one letter.`
	errorShortNameUnexpectedCharacter = `Invalid short name for option "%s".`
	errorLongNameUnexpectedCharacter = `Invalid long name for option "%s".`
	errorUnexpectedType = `Unepected type`
	errorInvalidValueBoolExpected = `Invalid value for option. Expected a value of type bool.`
	errorInvalidValueStringExpected = `Invalid value for option. Expected a value of type string.`

	// ----------------------------------------------------------------
	// spec.go
	// ----------------------------------------------------------------

	// TODO: unit test for errorInvalidCmdLineSpecInvalidOptionDefinition

	errorInvalidCmdLineSpecInvalidOptionDefinition = `Invalid command line specification. Invalid option found at position %d: "%s".`
	errorInvalidCmdLineSpecDuplicatedShortNamedOption = `Invalid command line specification. Duplicated name for short option at position %d: "%s".`
	errorInvalidCmdLineSpecDuplicatedLongNamedOption = `Invalid command line specification. Duplicated name for long option at positino %d: "%s".`
	errorInvalidCmdLineSpecReuseOfValueHolder = `Invalid command line specification. Duplicated value holder for the option at position %d.`
)


