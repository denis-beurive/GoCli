package cli

// This type defines the data structure used to organise the options definitions, relatively to their names.

type specIndex struct {
	Short map[string]*Option
	Long  map[string]*Option
}

// Return a pointer to the option's definition that applies to an option identified by its short name.
// If the given short name does not identify an option, then the function returns the value nil.

func (s *specIndex) getShortByName(inName string) *Option {
	if v, ok := s.Short[inName]; ok {
		return v
	}
	return nil
}

// Return a pointer to the option's definition that applies to an option identified by its long name.
// If the given long name does not identify an option, then the function returns the value nil.

func (s *specIndex) getLongByName(inName string) *Option {
	if v, ok := s.Long[inName]; ok {
		return v
	}
	return nil
}
