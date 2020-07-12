package flagenv

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

func ResetForTesting() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

const defaultOutput = `  -test_bool
    	bool value
`

func TestPrintDefaults(t *testing.T) {
	ResetForTesting()
	Bool("test_bool", "TEST_BOOL", false, "bool value")

	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)
	PrintDefaults()

	got := buf.String()
	if got != defaultOutput {
		t.Errorf("got %q want %q\n", got, defaultOutput)
	}
}

func TestFlags(t *testing.T) {
	ResetForTesting()

	// set some values
	os.Setenv("TEST_BOOL", "true")
	os.Setenv("TEST_INT", "1")
	os.Setenv("TEST_INT64", "2")
	os.Setenv("TEST_UINT", "3")
	os.Setenv("TEST_UINT64", "4")
	os.Setenv("TEST_FLOAT64", "5")
	os.Setenv("TEST_STRING", "text")

	// define flags
	Bool("test_bool", "TEST_BOOL", false, "bool value")
	String("test_string", "TEST_STRING", "", "string value")
	Int("test_int", "TEST_INT", 0, "int value")
	Int64("test_int64", "TEST_INT64", 0, "int64 value")
	Uint("test_uint", "TEST_UINT", 0, "uint value")
	Uint64("test_uint64", "TEST_UINT64", 0, "uint64 value")
	Float64("test_float64", "TEST_FLOAT64", 0, "float64 value")

	// get flag value and check if equal to setted env variables.
	visitor := func(f *flag.Flag) {
		if len(f.Name) > 5 && f.Name[0:5] == "test_" {
			g, ok := f.Value.(flag.Getter)
			if !ok {
				t.Errorf("Visit: value does not satisfy Getter: %T", f.Value)
				return
			}
			switch f.Name {
			case "test_bool":
				ok = g.Get() == true
			case "test_string":
				ok = g.Get() == "text"
			case "test_int":
				ok = g.Get() == int(1)
			case "test_int64":
				ok = g.Get() == int64(2)
			case "test_uint":
				ok = g.Get() == uint(3)
			case "test_uint64":
				ok = g.Get() == uint64(4)
			case "test_float64":
				ok = g.Get() == float64(5)
			}
			if !ok {
				t.Errorf("Visit: bad value %T(%v) for %s", g.Get(), g.Get(), f.Name)
			}
		}
	}
	flag.VisitAll(visitor)
}
