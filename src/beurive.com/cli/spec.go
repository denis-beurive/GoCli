package cli

import (
	"errors"
	"fmt"
)

// This structure represents the command line specification.

type Spec []Option


// Initialise the specification. This operation involves the following actions:
// - Check the specification (option names, holder types and holders singleness).
// - Build an index. Options are organised according to their length (short or long) and their names.

func (s Spec) init() (*specIndex, error) {
	shorts := make(map[string]*Option)
	longs := make(map[string]*Option)
	holders := make(map[interface{}]int)

	for i := range s {
		option := &s[i]
		if err := option.init(); nil != err {
			return nil, errors.New(fmt.Sprintf(errorInvalidCmdLineSpecInvalidOptionDefinition, i, err.Error()))
		}
		if "" != option.Short {
			if _, exists := shorts[option.Short]; exists {
				return nil, errors.New(fmt.Sprintf(errorInvalidCmdLineSpecDuplicatedShortNamedOption, i, option.Short))
			}
			shorts[option.Short] = option
		}
		if "" != option.Long {
			if _, exists := longs[option.Long]; exists {
				return nil, errors.New(fmt.Sprintf(errorInvalidCmdLineSpecDuplicatedLongNamedOption, i, option.Long))
			}
			longs[option.Long] = option
		}
		// Sanity check: make sure that the same variable is not used to store values for different options.
		if _, exists := holders[option.Holder]; ! exists {
			holders[option.Holder] = i
		} else {
			return nil, errors.New(fmt.Sprintf(errorInvalidCmdLineSpecReuseOfValueHolder, i))
		}

	}
	return &specIndex{Short: shorts, Long: longs }, nil
}

