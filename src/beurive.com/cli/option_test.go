package cli

import (
	"testing"
	"strings"
	"fmt"
	"strconv"
)

// -----------------------------------------------------------------
// Test the combination of short and long option names.
// -----------------------------------------------------------------

func TestInitNamesCombinationsOk(t *testing.T)  {
	var cloVerbose bool
	var o Option

	o = Option{Short: "v", Long: "verbose", Holder: &cloVerbose}
	if err := o.init(); nil != err {
		t.Error("Option's specifier should be valid!")
	}

	o = Option{Short: "", Long: "verbose", Holder: &cloVerbose}
	if err := o.init(); nil != err {
		t.Error("Option's specifier should be valid!")
	}

	o = Option{Short: "v", Long: "", Holder: &cloVerbose}
	if err := o.init(); nil != err {
		t.Error("Option's specifier should be valid!")
	}
}

// -----------------------------------------------------------------
// Test that the type of a value holder is well detected.
// -----------------------------------------------------------------

func TestGetTypeOk(t *testing.T)  {

	type testSet struct {
		Value    interface{}
		Expected typeOption
	}

	var vBool     bool
	var vString   string
	var vInt      int
	var vInt8     int8
	var vInt16    int16
	var vInt32    int32
	var vInt64    int64
	var vuInt     uint
	var vuInt8    uint8
	var vuInt16   uint16
	var vuInt32   uint32
	var vuInt64   uint64
	var vFloat32  float32
	var vFloat64  float64
	var vStrings  []string
	var vInts     []int
	var vInts8    []int8
	var vInts16   []int16
	var vInts32   []int32
	var vInts64   []int64
	var vuInts    []uint
	var vuInts8   []uint8
	var vuInts16  []uint16
	var vuInts32  []uint32
	var vuInts64  []uint64
	var vFloats32 []float32
	var vFloats64 []float64

	for i, v := range []testSet{
		testSet{ Value: &vBool, 	Expected: TypeBool },
		testSet{ Value: &vString, 	Expected: TypeString },
		testSet{ Value: &vInt, 		Expected: TypeInteger },
		testSet{ Value: &vInt8, 	Expected: TypeInteger8 },
		testSet{ Value: &vInt16, 	Expected: TypeInteger16 },
		testSet{ Value: &vInt32, 	Expected: TypeInteger32 },
		testSet{ Value: &vInt64, 	Expected: TypeInteger64 },
		testSet{ Value: &vuInt, 	Expected: TypeUInteger },
		testSet{ Value: &vuInt8, 	Expected: TypeUInteger8 },
		testSet{ Value: &vuInt16, 	Expected: TypeUInteger16 },
		testSet{ Value: &vuInt32, 	Expected: TypeUInteger32 },
		testSet{ Value: &vuInt64, 	Expected: TypeUInteger64 },
		testSet{ Value: &vFloat32, 	Expected: TypeFloat32 },
		testSet{ Value: &vFloat64, 	Expected: TypeFloat64 },
		testSet{ Value: &vStrings, 	Expected: TypeStrings },
		testSet{ Value: &vInts, 	Expected: TypeIntegers },
		testSet{ Value: &vInts8, 	Expected: TypeIntegers8 },
		testSet{ Value: &vInts16, 	Expected: TypeIntegers16 },
		testSet{ Value: &vInts32, 	Expected: TypeIntegers32 },
		testSet{ Value: &vInts64, 	Expected: TypeIntegers64 },
		testSet{ Value: &vuInts, 	Expected: TypeUIntegers },
		testSet{ Value: &vuInts8, 	Expected: TypeUIntegers8 },
		testSet{ Value: &vuInts16, 	Expected: TypeUIntegers16 },
		testSet{ Value: &vuInts32, 	Expected: TypeUIntegers32 },
		testSet{ Value: &vuInts64, 	Expected: TypeUIntegers64 },
		testSet{ Value: &vFloats32, 	Expected: TypeFloats32 },
		testSet{ Value: &vFloats64, 	Expected: TypeFloats64 },
	} {
		option := Option{ Short: "v", Long: "verbose", Holder: v.Value }
		if err := option.init(); nil != err {
			t.Errorf(`Unexpected error. Test #%d failed!`, i)
		} else {
			if ot, err := option.getType(); nil != err {
				t.Errorf(`Unexpected error. Test #%d failed!`, i)
			} else {
				if v.Expected != ot {
					t.Errorf(`Unexpected error. Test #%d failed!`, i)
				}
			}
		}
	}
}

// -----------------------------------------------------------------
// Test the types of the values holders.
// -----------------------------------------------------------------

func TestInitTypesOk(t *testing.T)  {
	var vBool     bool
	var vString   string
	var vInt      int
	var vInt8     int8
	var vInt16    int16
	var vInt32    int32
	var vInt64    int64
	var vuInt     uint
	var vuInt8    uint8
	var vuInt16   uint16
	var vuInt32   uint32
	var vuInt64   uint64
	var vFloat32  float32
	var vFloat64  float64
	var vStrings  []string
	var vInts     []int
	var vInts8    []int8
	var vInts16   []int16
	var vInts32   []int32
	var vInts64   []int64
	var vuInts    []uint
	var vuInts8   []uint8
	var vuInts16  []uint16
	var vuInts32  []uint32
	var vuInts64  []uint64
	var vFloats32 []float32
	var vFloats64 []float64

	for _, v := range []interface{}{
		&vBool,
		&vString,
		&vInt,
		&vInt8,
		&vInt16,
		&vInt32,
		&vInt64,
		&vuInt,
		&vuInt8,
		&vuInt16,
		&vuInt32,
		&vuInt64,
		&vFloat32,
		&vFloat64,
		&vStrings,
		&vInts,
		&vInts8,
		&vInts16,
		&vInts32,
		&vInts64,
		&vuInts,
		&vuInts8,
		&vuInts16,
		&vuInts32,
		&vuInts64,
		&vFloats32,
		&vFloats64} {
		o := Option{Short: "o", Long: "option", Holder: v}
		if err := o.init(); nil != err {
			t.Error(`Option's specifier should be valid!`)
		}
	}
}

// -----------------------------------------------------------------
// Test the addition of values into value holders.
// -----------------------------------------------------------------

func TestAddValueOk(t *testing.T)  {
	var vBool     bool
	var pBool     *bool
	var vString   string
	var pString   *string

	var vInt      int
	var pInt      *int
	var vInt8     int8
	var pInt8     *int8
	var vInt16    int16
	var pInt16    *int16
	var vInt32    int32
	var pInt32    *int32
	var vInt64    int64
	var pInt64    *int64
	var vuInt     uint
	var puInt     *uint
	var vuInt8    uint8
	var puInt8    *uint8
	var vuInt16   uint16
	var puInt16   *uint16
	var vuInt32   uint32
	var puInt32   *uint32
	var vuInt64   uint64
	var puInt64   *uint64
	var vFloat32  float32
	var pFloat32  *float32
	var vFloat64  float64
	var pFloat64  *float64
	var vStrings  []string = make([]string, 0)
	var pStrings  []string
	var vInts     []int = make([]int, 0)
	var pInts     []int
	var vInts8    []int8 = make([]int8, 0)
	var pInts8    []int8
	var vInts16   []int16 = make([]int16, 0)
	var pInts16   []int16
	var vInts32   []int32 = make([]int32, 0)
	var pInts32   []int32
	var vInts64   []int64 = make([]int64, 0)
	var pInts64   []int64
	var vuInts    []uint = make([]uint, 0)
	var puInts    []uint
	var vuInts8   []uint8 = make([]uint8, 0)
	var puInts8   []uint8
	var vuInts16  []uint16 = make([]uint16, 0)
	var puInts16  []uint16
	var vuInts32  []uint32 = make([]uint32, 0)
	var puInts32  []uint32
	var vuInts64  []uint64 = make([]uint64, 0)
	var puInts64  []uint64
	var vFloats32 []float32 = make([]float32, 0)
	var pFloats32 []float32
	var vFloats64 []float64 = make([]float64, 0)
	var pFloats64 []float64

	type setType struct {
		option *Option
		value interface{}
		expected interface{}
	}

	getOption := func(inShort string, inLong string, inHolder interface{}) *Option {
		return &Option{
			Short: inShort,
			Long: inLong,
			Holder: inHolder,
		}
	}

	testSet := []setType{
		{
			option: getOption("o", "--option", &vBool),
			value: true,
			expected: true,
		},
		{
			option: getOption("o", "--option", pBool),
			value: true,
			expected: true,
		},

		// --- Scalars ---

		{
			option: getOption("o", "--option", &vString),
			value: "/path/to/file",
			expected: "/path/to/file",
		},
		{
			option: getOption("o", "--option", pString),
			value: "/path/to/file",
			expected: "/path/to/file",
		},
		{
			option: getOption("o", "--option", &vInt),
			value: "10",
			expected: 10,
		},
		{
			option: getOption("o", "--option", pInt),
			value: "10",
			expected: 10,
		},
		{
			option: getOption("o", "--option", &vInt8),
			value: "10",
			expected: int8(10),
		},
		{
			option: getOption("o", "--option", pInt8),
			value: "10",
			expected: int8(10),
		},
		{
			option: getOption("o", "--option", &vInt16),
			value: "32767",
			expected: int16(32767),
		},
		{
			option: getOption("o", "--option", pInt16),
			value: "32767",
			expected: int16(32767),
		},
		{
			option: getOption("o", "--option", &vInt32),
			value: "2147483647",
			expected: int32(2147483647),
		},
		{
			option: getOption("o", "--option", pInt32),
			value: "2147483647",
			expected: int32(2147483647),
		},
		{
			option: getOption("o", "--option", &vInt64),
			value: "9223372036854775807",
			expected: int64(9223372036854775807),
		},
		{
			option: getOption("o", "--option", pInt64),
			value: "9223372036854775807",
			expected: int64(9223372036854775807),
		},
		{
			option: getOption("o", "--option", &vuInt),
			value: "255",
			expected: uint(255),
		},
		{
			option: getOption("o", "--option", puInt),
			value: "255",
			expected: uint(255),
		},
		{
			option: getOption("o", "--option", &vuInt8),
			value: "255",
			expected: uint8(255),
		},
		{
			option: getOption("o", "--option", puInt8),
			value: "255",
			expected: uint8(255),
		},
		{
			option: getOption("o", "--option", &vuInt16),
			value: "65535",
			expected: uint16(65535),
		},
		{
			option: getOption("o", "--option", puInt16),
			value: "65535",
			expected: uint16(65535),
		},
		{
			option: getOption("o", "--option", &vuInt32),
			value: "4294967295",
			expected: uint32(4294967295),
		},
		{
			option: getOption("o", "--option", puInt32),
			value: "4294967295",
			expected: uint32(4294967295),
		},
		{
			option: getOption("o", "--option", &vuInt64),
			value: "18446744073709551615",
			expected: uint64(18446744073709551615),
		},
		{
			option: getOption("o", "--option", puInt64),
			value: "18446744073709551615",
			expected: uint64(18446744073709551615),
		},
		{
			option: getOption("o", "--option", &vFloat32),
			value: strconv.FormatFloat(100000, 'f', -1, 32),
			expected: float32(100000),
		},
		{
			option: getOption("o", "--option", pFloat32),
			value: strconv.FormatFloat(100000, 'f', -1, 32),
			expected: float32(100000),
		},
		{
			option: getOption("o", "--option", &vFloat64),
			value: strconv.FormatFloat(100000, 'f', -1, 64),
			expected: float64(100000),
		},
		{
			option: getOption("o", "--option", pFloat64),
			value: strconv.FormatFloat(100000, 'f', -1, 64),
			expected: float64(100000),
		},

		// --- Arrays ---

		{
			option: getOption("o", "--option", &vStrings),
			value: "/path/to/file",
			expected: []string{"/path/to/file"},
		},
		{
			option: getOption("o", "--option", &pStrings),
			value: "/path/to/file",
			expected: []string{"/path/to/file"},
		},
		{
			option: getOption("o", "--option", &vInts),
			value: "10",
			expected: []int{10},
		},
		{
			option: getOption("o", "--option", &pInts),
			value: "10",
			expected: []int{10},
		},
		{
			option: getOption("o", "--option", &vInts8),
			value: "10",
			expected: []int8{10},
		},
		{
			option: getOption("o", "--option", &pInts8),
			value: "10",
			expected: []int8{10},
		},
		{
			option: getOption("o", "--option", &vInts16),
			value: "32767",
			expected: []int16{32767},
		},
		{
			option: getOption("o", "--option", &pInts16),
			value: "32767",
			expected: []int16{32767},
		},
		{
			option: getOption("o", "--option", &vInts32),
			value: "2147483647",
			expected: []int32{2147483647},
		},
		{
			option: getOption("o", "--option", &pInts32),
			value: "2147483647",
			expected: []int32{2147483647},
		},
		{
			option: getOption("o", "--option", &vInts64),
			value: "9223372036854775807",
			expected: []int64{9223372036854775807},
		},
		{
			option: getOption("o", "--option", &pInts64),
			value: "9223372036854775807",
			expected: []int64{9223372036854775807},
		},
		{
			option: getOption("o", "--option", &vuInts),
			value: "10",
			expected: []uint{10},
		},
		{
			option: getOption("o", "--option", &puInts),
			value: "10",
			expected: []uint{10},
		},
		{
			option: getOption("o", "--option", &vuInts8),
			value: "255",
			expected: []uint8{255},
		},
		{
			option: getOption("o", "--option", &puInts8),
			value: "255",
			expected: []uint8{255},
		},
		{
			option: getOption("o", "--option", &vuInts16),
			value: "65535",
			expected: []uint16{65535},
		},
		{
			option: getOption("o", "--option", &puInts16),
			value: "65535",
			expected: []uint16{65535},
		},
		{
			option: getOption("o", "--option", &vuInts32),
			value: "4294967295",
			expected: []uint32{4294967295},
		},
		{
			option: getOption("o", "--option", &puInts32),
			value: "4294967295",
			expected: []uint32{4294967295},
		},
		{
			option: getOption("o", "--option", &vuInts64),
			value: "9223372036854775807",
			expected: []uint64{9223372036854775807},
		},
		{
			option: getOption("o", "--option", &puInts64),
			value: "9223372036854775807",
			expected: []uint64{9223372036854775807},
		},
		{
			option: getOption("o", "--option", &vFloats32),
			value: strconv.FormatFloat(100000, 'f', -1, 32),
			expected: []float32{100000},
		},
		{
			option: getOption("o", "--option", &pFloats32),
			value: strconv.FormatFloat(100000, 'f', -1, 32),
			expected: []float32{100000},
		},
		{
			option: getOption("o", "--option", &vFloats64),
			value: strconv.FormatFloat(100000, 'f', -1, 64),
			expected: []float64{100000},
		},
		{
			option: getOption("o", "--option", &pFloats64),
			value: strconv.FormatFloat(100000, 'f', -1, 64),
			expected: []float64{100000},
		},
	}

	for i, set := range testSet {
		if err := set.option.addValue(set.value); nil != err {
			t.Errorf(`Test #%d failed. The test should be OK. %s`, i, err.Error())
		} else {
			if set.option.isSingleton() {
				tp, _ := set.option.getType()
				if TypeBool == tp {
					if v, ok := set.option.Holder.(*bool); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "bool".`, i)
					} else {
						ev, _ := set.expected.(bool)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeString == tp {
					if v, ok := set.option.Holder.(*string); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*string".`, i)
					} else {
						ev, _ := set.expected.(string)
						if 0 != strings.Compare(*v, ev) {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeInteger == tp {
					if v, ok := set.option.Holder.(*int); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*int".`, i)
					} else {
						ev, _ := set.expected.(int)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeInteger8 == tp {
					if v, ok := set.option.Holder.(*int8); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*int8".`, i)
					} else {
						ev, _ := set.expected.(int8);
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value! Expected %d, got %d`, i, ev, *v)
						}
					}
				} else if TypeInteger16 == tp {
					if v, ok := set.option.Holder.(*int16); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*int16".`, i)
					} else {
						ev, _ := set.expected.(int16)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeInteger32 == tp {
					if v, ok := set.option.Holder.(*int32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*int32".`, i)
					} else {
						ev, _ := set.expected.(int32)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeInteger64 == tp {
					if v, ok := set.option.Holder.(*int64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*int64".`, i)
					} else {
						ev, _ := set.expected.(int64)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUInteger == tp {
					if v, ok := set.option.Holder.(*uint); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*uint".`, i)
					} else {
						ev, _ := set.expected.(uint)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUInteger8 == tp {
					if v, ok := set.option.Holder.(*uint8); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*uint8".`, i)
					} else {
						ev, _ := set.expected.(uint8);
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value! Expected %d, got %d`, i, ev, *v)
						}
					}
				} else if TypeUInteger16 == tp {
					if v, ok := set.option.Holder.(*uint16); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*uint16".`, i)
					} else {
						ev, _ := set.expected.(uint16)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUInteger32 == tp {
					if v, ok := set.option.Holder.(*uint32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*uint32".`, i)
					} else {
						ev, _ := set.expected.(uint32)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUInteger64 == tp {
					if v, ok := set.option.Holder.(*uint64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*uint64".`, i)
					} else {
						ev, _ := set.expected.(uint64)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloat32 == tp {
					if v, ok := set.option.Holder.(*float32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*float32".`, i)
					} else {
						ev, _ := set.expected.(float32)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloat64 == tp {
					if v, ok := set.option.Holder.(*float64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*float64".`, i)
					} else {
						ev, _ := set.expected.(float64)
						if *v != ev {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else {
					t.Errorf(`Test #%d failed. Unexpected type encountered!`, i)
				}
			} else if set.option.requireValue() {
				tp, _ := set.option.getType()
				if TypeStrings == tp {
					if v, ok := set.option.Holder.(*[]string); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]string".`, i)
					} else {
						ev, _ := set.expected.([]string)
						if 0 != strings.Compare((*v)[0], ev[0]) {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeIntegers == tp {
					if v, ok := set.option.Holder.(*[]int); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]int".`, i)
					} else {
						ev, _ := set.expected.([]int)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeIntegers8 == tp {
					if v, ok := set.option.Holder.(*[]int8); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]int8".`, i)
					} else {
						ev, _ := set.expected.([]int8);
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value! Expected %d, got %d`, i, ev[0], (*v)[0])
						}
					}
				} else if TypeIntegers16 == tp {
					if v, ok := set.option.Holder.(*[]int16); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]int16".`, i)
					} else {
						ev, _ := set.expected.([]int16)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeIntegers32 == tp {
					if v, ok := set.option.Holder.(*[]int32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]int32".`, i)
					} else {
						ev, _ := set.expected.([]int32)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeIntegers64 == tp {
					if v, ok := set.option.Holder.(*[]int64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]int64".`, i)
					} else {
						ev, _ := set.expected.([]int64)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloats32 == tp {
					if v, ok := set.option.Holder.(*[]float32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]float32".`, i)
					} else {
						ev, _ := set.expected.([]float32)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloats64 == tp {
					if v, ok := set.option.Holder.(*[]float64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]float64".`, i)
					} else {
						ev, _ := set.expected.([]float64)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}					}
				} else if TypeUIntegers == tp {
					if v, ok := set.option.Holder.(*[]uint); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]uint".`, i)
					} else {
						ev, _ := set.expected.([]uint)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUIntegers8 == tp {
					if v, ok := set.option.Holder.(*[]uint8); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]uint8".`, i)
					} else {
						ev, _ := set.expected.([]uint8);
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value! Expected %d, got %d`, i, ev[0], (*v)[0])
						}
					}
				} else if TypeUIntegers16 == tp {
					if v, ok := set.option.Holder.(*[]uint16); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]uint16".`, i)
					} else {
						ev, _ := set.expected.([]uint16)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUIntegers32 == tp {
					if v, ok := set.option.Holder.(*[]uint32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]uint32".`, i)
					} else {
						ev, _ := set.expected.([]uint32)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeUIntegers64 == tp {
					if v, ok := set.option.Holder.(*[]uint64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]uint64".`, i)
					} else {
						ev, _ := set.expected.([]uint64)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloats32 == tp {
					if v, ok := set.option.Holder.(*[]float32); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]float32".`, i)
					} else {
						ev, _ := set.expected.([]float32)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}
					}
				} else if TypeFloats64 == tp {
					if v, ok := set.option.Holder.(*[]float64); ! ok {
						t.Errorf(`Test #%d failed. The expected type of the value should be "*[]float64".`, i)
					} else {
						ev, _ := set.expected.([]float64)
						if (*v)[0] != ev[0] {
							t.Errorf(`Test #%d failed. Invalid value!`, i)
						}					}
				} else {
					t.Errorf(`Test #%d failed. Unexpected type encountered! %d`, i, tp)
				}
			}
		}
	}

}

// -----------------------------------------------------------------
// Test option's constraints "singleton" / "require value".
// -----------------------------------------------------------------

func TestOptionConstraints(t *testing.T)  {
	var vBool     bool
	var vString   string
	var vInt      int
	var vInt8     int8
	var vInt16    int16
	var vInt32    int32
	var vInt64    int64
	var vuInt     uint
	var vuInt8    uint8
	var vuInt16   uint16
	var vuInt32   uint32
	var vuInt64   uint64
	var vFloat32  float32
	var vFloat64  float64
	var vStrings  []string
	var vInts     []int
	var vInts8    []int8
	var vInts16   []int16
	var vInts32   []int32
	var vInts64   []int64
	var vuInts    []uint
	var vuInts8   []uint8
	var vuInts16  []uint16
	var vuInts32  []uint32
	var vuInts64  []uint64
	var vFloats32 []float32
	var vFloats64 []float64

	var bool2string = func (b bool) string {
		if b { return "true"}
		return "false"
	}

	type data struct {
		holder interface{}
		singleton bool
		require bool
	}

	for i, v := range []data{
		{ holder: &vBool,     singleton: true,   require: false },
		{ holder: &vString,   singleton: true,   require: true },
		{ holder: &vInt,      singleton: true,   require: true },
		{ holder: &vInt8,     singleton: true,   require: true },
		{ holder: &vInt16,    singleton: true,   require: true },
		{ holder: &vInt32,    singleton: true,   require: true },
		{ holder: &vInt64,    singleton: true,   require: true },
		{ holder: &vuInt,     singleton: true,   require: true },
		{ holder: &vuInt8,    singleton: true,   require: true },
		{ holder: &vuInt16,   singleton: true,   require: true },
		{ holder: &vuInt32,   singleton: true,   require: true },
		{ holder: &vuInt64,   singleton: true,   require: true },
		{ holder: &vFloat32,  singleton: true,   require: true },
		{ holder: &vFloat64,  singleton: true,   require: true },
		{ holder: &vStrings,  singleton: false,  require: true },
		{ holder: &vInts,     singleton: false,  require: true },
		{ holder: &vInts8,    singleton: false,  require: true },
		{ holder: &vInts16,   singleton: false,  require: true },
		{ holder: &vInts32,   singleton: false,  require: true },
		{ holder: &vInts64,   singleton: false,  require: true },
		{ holder: &vuInts,    singleton: false,  require: true },
		{ holder: &vuInts8,   singleton: false,  require: true },
		{ holder: &vuInts16,  singleton: false,  require: true },
		{ holder: &vuInts32,  singleton: false,  require: true },
		{ holder: &vuInts64,  singleton: false,  require: true },
		{ holder: &vFloats32, singleton: false,  require: true },
		{ holder: &vFloats64, singleton: false,  require: true }} {
		o := Option{Short: "o", Long: "option", Holder: v.holder}
		if o.isSingleton() != v.singleton {
			t.Errorf("%d: %s != %s", i, bool2string(o.isSingleton()), bool2string(v.singleton) )
		}
		if o.requireValue() != v.require {
			t.Errorf("%d: %s != %s", i, bool2string(o.requireValue()), bool2string(v.require) )
		}
	}
}

// -----------------------------------------------------------------
// Test the error messages.
// -----------------------------------------------------------------

func TestEM_InvalidOptionSpecificationNoName(t *testing.T)  {
	var cloVerbose bool
	var o Option

	o = Option{Short: "", Long: "", Holder: &cloVerbose}
	if err := o.init(); nil == err {
		t.Error("Option's specifier should not be valid!")
	} else {
		if 0 != strings.Compare(err.Error(), errorInvalidOptionSpecificationNoName) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), errorInvalidOptionSpecificationNoName)
		}
	}
}

func TestEM_ShortNameTooLong(t *testing.T)  {
	var cloVerbose bool
	var o Option

	o = Option{Short: "vv", Long: "", Holder: &cloVerbose}
	if err := o.init(); nil == err {
		t.Error("Option's specifier should not be valid!")
	} else {
		m := fmt.Sprintf(errorShortNameTooLong, "vv")
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

func TestEM_ShortNameUnexpectedCharacter(t *testing.T)  {
	var cloVerbose bool
	var o Option

	o = Option{Short: "#", Long: "", Holder: &cloVerbose}
	if err := o.init(); nil == err {
		t.Error("Option's specifier should not be valid!")
	} else {
		m := fmt.Sprintf(errorShortNameUnexpectedCharacter, "#")
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

func TestEM_LongNameUnexpectedCharacter(t *testing.T)  {

	var cloVerbose bool
	var o Option


	o = Option{Short: "", Long: "v#", Holder: &cloVerbose}
	if err := o.init(); nil == err {
		t.Error("Option's specifier should not be valid!")
	} else {
		m := fmt.Sprintf(errorLongNameUnexpectedCharacter, "v#")
		if 0 != strings.Compare(err.Error(), m) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), m)
		}
	}
}

func TestEM_InvalidOptionSpecificationUnexpectedHolderType(t *testing.T)  {
	var cloError []bool
	var o Option

	o = Option{Short: "v", Long: "", Holder: &cloError}
	if err := o.init(); nil == err {
		t.Error("Option's specifier should not be valid!")
	} else {
		if 0 != strings.Compare(err.Error(), errorInvalidOptionSpecificationUnexpectedHolderType) {
			t.Errorf(`Test failed. Got [%s] / [%s]`, err.Error(), errorInvalidOptionSpecificationUnexpectedHolderType)
		}
	}
}

func TestEM_InvalidValueBoolExpected(t *testing.T)  {
	// The value holder is a boolean.
	// Therefore, the value given to the init() function must be a boolean.

	var cloVerbose bool

	o := Option{ Short:"v", Long:"verbose", Holder:&cloVerbose }
	if err := o.init(); nil != err {
		t.Error(`Unexpected error: the initialisation process should be OK.`)
	}

	if err := o.addValue(10); nil == err {
		t.Error(`Unexpected error! The test should fail!`)
	} else {
		if 0 != strings.Compare(errorInvalidValueBoolExpected, err.Error()) {
			t.Errorf(`Unexpected error! Expected "%s", got "%s".`,errorInvalidValueBoolExpected, err.Error())
		}
	}
}

func TestEM_InvalidValueStringExpected(t *testing.T)  {
	// The value holder is not a boolean.
	// Therefore, the value given to the init() function must be a string.

	var cloPath string

	o := Option{ Short:"p", Long:"path", Holder:&cloPath }
	if err := o.init(); nil != err {
		t.Error(`Unexpected error: the initialisation process should be OK.`)
	}

	if err := o.addValue(10); nil == err {
		t.Error(`Unexpected error! The test should fail!`)
	} else {
		if 0 != strings.Compare(errorInvalidValueStringExpected, err.Error()) {
			t.Errorf(`Unexpected error! Expected "%s", got "%s".`,errorInvalidValueStringExpected, err.Error())
		}
	}
}

func TestEM_UnexpectedType(t *testing.T) {
	var v typeOption
	option := Option{ Short:"v", Long:"verbose", Holder: &v}
	if _, err := option.getType(); nil == err {
		t.Error(`The test should fail!`)
	} else {
		if 0 != strings.Compare(errorUnexpectedType, err.Error()) {
			t.Errorf(`Invalid error message. Expected "%s", got "%s" !`, errorUnexpectedType, err.Error())
		}
	}
}




