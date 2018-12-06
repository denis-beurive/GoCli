package cli

import (
	"errors"
	"fmt"
	"strconv"
)

// The type Option represents an option within the command line.
// An option is represented by the elements listed below:
//
// - an optional short name (or nickname).
// - an optional long name (or full name).
// - a mandatory "value holder". The value holder is a pointer to a variable used to store the value defined by the option.
//
// Within the command line, short names are prefixed by a single dash (ex: -i).
// Within the command line, long names are prefixed by a double dash (ex: --input).
// Value holders points to variables which types may be:
//
// - bool: this type is used to store the "state" (set or unset) of an option that does not require a value (a flag or
//   a switch).
// - string: this type is used to store the value of an option that defines a string.
// - int: this type is used to store the value of an option that defines an integer.
// - []string: this type is used to store a list of strings. Each element of the list comes from one occurrence of the
//   option's name, within the command line (ex: "--path=/usr/bin --path=/bin" will be stored as "[]string{`/usr/bin`, `/bin`}").
// - []int: this type is used to store a list of integers. Each element of the list comes from one occurrence of the
//   option's name, within the command line (ex: "--level=0 --level=1" will be stored as "[]int{0, 1}").

// This type represents the "GO type" of the option's value holder.

type typeOption int

const (
	TypeUnexpected typeOption = iota
	TypeBool
	TypeString
	TypeInteger
	TypeInteger8
	TypeInteger16
	TypeInteger32
	TypeInteger64
	TypeFloat32
	TypeFloat64
	TypeStrings
	TypeIntegers
	TypeIntegers8
	TypeIntegers16
	TypeIntegers32
	TypeIntegers64
	TypeFloats32
	TypeFloats64

	TypeUInteger
	TypeUInteger8
	TypeUInteger16
	TypeUInteger32
	TypeUInteger64

	TypeUIntegers
	TypeUIntegers8
	TypeUIntegers16
	TypeUIntegers32
	TypeUIntegers64
)

// This type defines the constraints that apply to a type of option's value holder.

type typeConstraints struct {
	singleton bool
	value bool
}

// For all "GO types" that can be used as options' holders, this map defines the following constraints:
// - Can the option appears more than once within the command line ?
// - Does the option accept value(s) ?

var typesConstraints = map[typeOption]typeConstraints{
	TypeBool:        {singleton:true,  value:false},
	TypeString:      {singleton:true,  value:true},
	TypeInteger:     {singleton:true,  value:true},
	TypeInteger8:    {singleton:true,  value:true},
	TypeInteger16:   {singleton:true,  value:true},
	TypeInteger32:   {singleton:true,  value:true},
	TypeInteger64:   {singleton:true,  value:true},
	TypeFloat32:     {singleton:true,  value:true},
	TypeFloat64:     {singleton:true,  value:true},
	TypeStrings:     {singleton:false, value:true},
	TypeIntegers:    {singleton:false, value:true},
	TypeIntegers8:   {singleton:false, value:true},
	TypeIntegers16:  {singleton:false, value:true},
	TypeIntegers32:  {singleton:false, value:true},
	TypeIntegers64:  {singleton:false, value:true},
	TypeFloats32:    {singleton:false, value:true},
	TypeFloats64:    {singleton:false, value:true},

	TypeUInteger:    {singleton:true,  value:true},
	TypeUInteger8:   {singleton:true,  value:true},
	TypeUInteger16:  {singleton:true,  value:true},
	TypeUInteger32:  {singleton:true,  value:true},
	TypeUInteger64:  {singleton:true,  value:true},

	TypeUIntegers:   {singleton:false, value:true},
	TypeUIntegers8:  {singleton:false, value:true},
	TypeUIntegers16: {singleton:false, value:true},
	TypeUIntegers32: {singleton:false, value:true},
	TypeUIntegers64: {singleton:false, value:true},
}

// This structure defines an option.
// * The attribute "Short" represents the short name of the option. The short name is made of one, and only one, character.
// * The attribute "Long" represents the long name of the option. The long name is made of one or more characters.
// * The attribute "Holder" contains a pointer to the required data type. Please note that this pointer may point to an
//   allocated variable or not. In the latter case, the variable will be allocated for you.
// * The attribute "set" indicates whether the option is set or not.
//   The value true indicates that the option is set.
//   The value false indicates that the option is not set.

type Option struct {
	Short string        // The option's short name.
	Long string         // The option's long name.
	Holder interface{}  // pointer to the option's value holder.
	set bool            // The flag that specifies whether the option is set or not.
}

// Test whether an option is set or not.
// If the option is set, then the function returns the value true.
// Otherwise, it returns the value false.

func (o *Option) isSet() bool {
	return o.set
}

// Declare an option as being set.

func (o *Option) setIt() {
	o.set = true
}

// Add a value to an option.
// Please note that the parameter inValue may be a string or a boolean.

func (o *Option) addValue(inValue interface{}) error {

	typeOption, _ := o.getType()

	if TypeBool == typeOption {
		if v, ok := inValue.(bool); ! ok {
			return errors.New(errorInvalidValueBoolExpected)
		} else {
			p, _ := o.Holder.(*bool)
			if nil == p {
				o.Holder = new(bool)
				p, _ = o.Holder.(*bool)
			}
			*p = v
		}
		return nil
	}

	v, ok := inValue.(string);
	if ! ok {
		return errors.New(errorInvalidValueStringExpected)
	}

	switch typeOption {
		case TypeString:
			p, _ := o.Holder.(*string)
			if nil == p {
				o.Holder = new(string)
				p, _ = o.Holder.(*string)
			}
			*p = v
		case TypeInteger:
			// Be careful, the variable "p" is a copy of the variable "o.Holder".
			// Affecting a value to "p" does not modify the value of "o.Holder".
			p, _ := o.Holder.(*int)

			// Be careful, "o.Holder" is an interface.
			// We cannot directly compare an interface's against nil.
			if nil == p {
				o.Holder = new(int)
				p, _ = o.Holder.(*int) // The value of "o.Holder" changed!
			}
			v, err := strconv.ParseInt(v, 10, 0)
			if nil != err { return err }
			*p = int(v)
		case TypeInteger8:
			p, _ := o.Holder.(*int8)
			if nil == p {
				o.Holder = new(int8)
				p, _ = o.Holder.(*int8)
			}
			v, err := strconv.ParseInt(v, 10, 8)
			if nil != err { return err }
			*p = int8(v)
		case TypeInteger16:
			p, _ := o.Holder.(*int16)
			if nil == p {
				o.Holder = new(int16)
				p, _ = o.Holder.(*int16)
			}
			v, err := strconv.ParseInt(v, 10, 16)
			if nil != err { return err }
			*p = int16(v)
		case TypeInteger32:
			p, _ := o.Holder.(*int32)
			if nil == p {
				o.Holder = new(int32)
				p, _ = o.Holder.(*int32)
			}
			v, err := strconv.ParseInt(v, 10, 32)
			if nil != err { return err }
			*p = int32(v)
		case TypeInteger64:
			p, _ := o.Holder.(*int64)
			if nil == p {
				o.Holder = new(int64)
				p, _ = o.Holder.(*int64)
			}
			v, err := strconv.ParseInt(v, 10, 64)
			if nil != err { return err }
			*p = v
		case TypeUInteger:
			p, _ := o.Holder.(*uint)
			if nil == p {
				o.Holder = new(uint)
				p, _ = o.Holder.(*uint) // The value of "o.Holder" changed!
			}
			v, err := strconv.ParseUint(v, 10, 0)
			if nil != err { return err }
			*p = uint(v)
		case TypeUInteger8:
			p, _ := o.Holder.(*uint8)
			if nil == p {
				o.Holder = new(uint8)
				p, _ = o.Holder.(*uint8)
			}
			v, err := strconv.ParseUint(v, 10, 8)
			if nil != err { return err }
			*p = uint8(v)
		case TypeUInteger16:
			p, _ := o.Holder.(*uint16)
			if nil == p {
				o.Holder = new(uint16)
				p, _ = o.Holder.(*uint16)
			}
			v, err := strconv.ParseUint(v, 10, 16)
			if nil != err { return err }
			*p = uint16(v)
		case TypeUInteger32:
			p, _ := o.Holder.(*uint32)
			if nil == p {
				o.Holder = new(uint32)
				p, _ = o.Holder.(*uint32)
			}
			v, err := strconv.ParseUint(v, 10, 32)
			if nil != err { return err }
			*p = uint32(v)
		case TypeUInteger64:
			p, _ := o.Holder.(*uint64)
			if nil == p {
				o.Holder = new(uint64)
				p, _ = o.Holder.(*uint64)
			}
			v, err := strconv.ParseUint(v, 10, 64)
			if nil != err { return err }
			*p = v
		case TypeFloat32:
			p, _ := o.Holder.(*float32)
			if nil == p {
				o.Holder = new(float32)
				p, _ = o.Holder.(*float32)
			}
			v, err := strconv.ParseFloat(v, 32)
			if nil != err { return err }
			*p = float32(v)
		case TypeFloat64:
			p, _ := o.Holder.(*float64)
			if nil == p {
				o.Holder = new(float64)
				p, _ = o.Holder.(*float64)
			}
			v, err := strconv.ParseFloat(v, 64)
			if nil != err { return err }
			*p = v
		case TypeStrings:
			p, _ := o.Holder.(*[]string)
			if nil == p { *p = make([]string, 0) }
			*p = append(*p, v)


		case TypeIntegers:
			v, err := strconv.ParseInt(v, 10, 0)
			if nil != err { return err }
			p, _ := o.Holder.(*[]int)
			if nil == p { *p = make([]int, 0) }
			*p = append(*p, int(v))
		case TypeIntegers8:
			v, err := strconv.ParseInt(v, 10, 8)
			if nil != err { return err }
			p, _ := o.Holder.(*[]int8)
			if nil == p { *p = make([]int8, 0) }
			*p = append(*p, int8(v))
		case TypeIntegers16:
			v, err := strconv.ParseInt(v, 10, 16)
			if nil != err { return err }
			p, _ := o.Holder.(*[]int16)
			if nil == p { *p = make([]int16, 0) }
			*p = append(*p, int16(v))
		case TypeIntegers32:
			v, err := strconv.ParseInt(v, 10, 32)
			if nil != err { return err }
			p, _ := o.Holder.(*[]int32)
			if nil == p { *p = make([]int32, 0) }
			*p = append(*p, int32(v))
		case TypeIntegers64:
			v, err := strconv.ParseInt(v, 10, 64)
			if nil != err { return err }
			p, _ := o.Holder.(*[]int64)
			if nil == p { *p = make([]int64, 0) }
			*p = append(*p, v)
		case TypeUIntegers:
			v, err := strconv.ParseUint(v, 10, 0)
			if nil != err { return err }
			p, _ := o.Holder.(*[]uint)
			if nil == p { *p = make([]uint, 0) }
			*p = append(*p, uint(v))
		case TypeUIntegers8:
			v, err := strconv.ParseUint(v, 10, 8)
			if nil != err { return err }
			p, _ := o.Holder.(*[]uint8)
			if nil == p { *p = make([]uint8, 0) }
			*p = append(*p, uint8(v))
		case TypeUIntegers16:
			v, err := strconv.ParseUint(v, 10, 16)
			if nil != err { return err }
			p, _ := o.Holder.(*[]uint16)
			if nil == p { *p = make([]uint16, 0) }
			*p = append(*p, uint16(v))
		case TypeUIntegers32:
			v, err := strconv.ParseUint(v, 10, 32)
			if nil != err { return err }
			p, _ := o.Holder.(*[]uint32)
			if nil == p { *p = make([]uint32, 0) }
			*p = append(*p, uint32(v))
		case TypeUIntegers64:
			v, err := strconv.ParseUint(v, 10, 64)
			if nil != err { return err }
			p, _ := o.Holder.(*[]uint64)
			if nil == p { *p = make([]uint64, 0) }
			*p = append(*p, v)
		case TypeFloats32:
			v, err := strconv.ParseFloat(v, 32)
			if nil != err { return err }
			p, _ := o.Holder.(*[]float32)
			if nil == p { *p = make([]float32, 0) }
			*p = append(*p, float32(v))
		case TypeFloats64:
			v, err := strconv.ParseFloat(v, 64)
			if nil != err { return err }
			p, _ := o.Holder.(*[]float64)
			if nil == p { *p = make([]float64, 0) }
			*p = append(*p, v)
	}
	return nil
}

// Test whether an option can appear only once within the command line or not.

func (o *Option) isSingleton() bool {
	t, _ := o.getType();
	return typesConstraints[t].singleton
}

// Test whether an option requires a value or not.

func (o *Option) requireValue() bool {
	t, _ := o.getType();
	return typesConstraints[t].value
}

// Return the long name associated to an option, if it exists.

func (o *Option) getLong() (string, bool) {
	if "" != o.Long { return o.Long, true }
	return "", false
}

// Return the short name associated to an option, if it exists.

func (o *Option) getShort() (string, bool) {
	if "" != o.Short { return o.Short, true }
	return "", false
}

// Initialise an option. The initialisation consists of the actions listed below:
// - Checks that at least one name is specified (the short one or the long one).
// - Checks that the type of the variable used to store the option's value is valid.
// - Initialises the value of the option, in the case of an option that does not take values.
// - Initialises the state (set or unset) of the option. the state is initialised to the value false (unset).

func (o Option) init() error {

	if "" == o.Short && "" == o.Long {
		return errors.New(errorInvalidOptionSpecificationNoName)
	}
	if "" != o.Short {
		if len(o.Short) > 1 {
			return errors.New(fmt.Sprintf(errorShortNameTooLong, o.Short))
		}
		if ok, _ := isOptionShort(o.Short); ! ok {
			return errors.New(fmt.Sprintf(errorShortNameUnexpectedCharacter, o.Short))
		}
	}
	if "" != o.Long {
		if ok, _ := isOptionLong(o.Long); ! ok {
			return errors.New(fmt.Sprintf(errorLongNameUnexpectedCharacter, o.Long))
		}
	}
	if _, err := o.getType(); nil != err {
		return errors.New(errorInvalidOptionSpecificationUnexpectedHolderType)
	}

	o.set = false
	if t, _ := o.getType(); TypeBool == t {
		p, _ := o.Holder.(*bool)
		*p = false
	}

	return nil
}

// Return the type of the variable used to store the option's value.
// This type defines the constraints that apply to the option's value.

func (o Option) getType() (typeOption, error) {
	if _, ok := o.Holder.(*bool); ok { return TypeBool, nil }

	// Single value
	if _, ok := o.Holder.(*string);    ok { return TypeString,      nil }
	if _, ok := o.Holder.(*int);       ok { return TypeInteger,     nil }
	if _, ok := o.Holder.(*int8);      ok { return TypeInteger8,    nil }
	if _, ok := o.Holder.(*int16);     ok { return TypeInteger16,   nil }
	if _, ok := o.Holder.(*int32);     ok { return TypeInteger32,   nil }
	if _, ok := o.Holder.(*int64);     ok { return TypeInteger64,   nil }
	if _, ok := o.Holder.(*float32);   ok { return TypeFloat32,     nil }
	if _, ok := o.Holder.(*float64);   ok { return TypeFloat64,     nil }
	if _, ok := o.Holder.(*uint);      ok { return TypeUInteger,    nil }
	if _, ok := o.Holder.(*uint8);     ok { return TypeUInteger8,   nil }
	if _, ok := o.Holder.(*uint16);    ok { return TypeUInteger16,  nil }
	if _, ok := o.Holder.(*uint32);    ok { return TypeUInteger32,  nil }
	if _, ok := o.Holder.(*uint64);    ok { return TypeUInteger64,  nil }

	// Multiple values
	if _, ok := o.Holder.(*[]string);  ok { return TypeStrings,     nil }
	if _, ok := o.Holder.(*[]int);     ok { return TypeIntegers,    nil }
	if _, ok := o.Holder.(*[]int8);    ok { return TypeIntegers8,   nil }
	if _, ok := o.Holder.(*[]int16);   ok { return TypeIntegers16,  nil }
	if _, ok := o.Holder.(*[]int32);   ok { return TypeIntegers32,  nil }
	if _, ok := o.Holder.(*[]int64);   ok { return TypeIntegers64,  nil }
	if _, ok := o.Holder.(*[]float32); ok { return TypeFloats32,    nil }
	if _, ok := o.Holder.(*[]float64); ok { return TypeFloats64,    nil }
	if _, ok := o.Holder.(*[]uint);    ok { return TypeUIntegers,   nil }
	if _, ok := o.Holder.(*[]uint8);   ok { return TypeUIntegers8,  nil }
	if _, ok := o.Holder.(*[]uint16);  ok { return TypeUIntegers16, nil }
	if _, ok := o.Holder.(*[]uint32);  ok { return TypeUIntegers32, nil }
	if _, ok := o.Holder.(*[]uint64);  ok { return TypeUIntegers64, nil }

	return TypeUnexpected, errors.New(errorUnexpectedType)
}

