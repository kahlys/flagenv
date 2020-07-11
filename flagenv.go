package flagenv

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Parse parses the command-line/environment-variable flags.
func Parse() {
	flag.Parse()
}

// PrintDefaults prints a usage message showing the default settings of all defined command-line flags.
func PrintDefaults() {
	flag.PrintDefaults()
}

// String defines a string flag with specified name, env variable, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, env string, value string, usage string) *string {
	return flag.String(name, flagEnvString(env, value), usage)
}

func flagEnvString(key string, value string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return value
}

// Bool defines a bool flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, env string, value bool, usage string) *bool {
	return flag.Bool(name, flagEnvBool(env, value), usage)
}

func flagEnvBool(key string, value bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : boolean parsing error`, val, key))
		}
		return v
	}
	return value
}

// Int defines a int flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Int(name string, env string, value int, usage string) *int {
	return flag.Int(name, flagEnvInt(env, value), usage)
}

func flagEnvInt(key string, value int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : integer parsing error`, val, key))
		}
		return v
	}
	return value
}

// Int64 defines a int64 flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Int64(name string, env string, value int64, usage string) *int64 {
	return flag.Int64(name, flagEnvInt64(env, value), usage)
}

func flagEnvInt64(key string, value int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : integer parsing error`, val, key))
		}
		return v
	}
	return value
}

// Uint defines a uint flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Uint(name string, env string, value uint, usage string) *uint {
	return flag.Uint(name, flagEnvUint(env, value), usage)
}

func flagEnvUint(key string, value uint) uint {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : integer parsing error`, val, key))
		}
		return uint(v)
	}
	return value
}

// Uint64 defines a uint64 flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Uint64(name string, env string, value uint64, usage string) *uint64 {
	return flag.Uint64(name, flagEnvUint64(env, value), usage)
}

func flagEnvUint64(key string, value uint64) uint64 {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : integer parsing error`, val, key))
		}
		return v
	}
	return value
}

// Float64 defines a float64 flag with specified name, env variable, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Float64(name string, env string, value float64, usage string) *float64 {
	return flag.Float64(name, flagEnvFloat64(env, value), usage)
}

func flagEnvFloat64(key string, value float64) float64 {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			outputErr(fmt.Sprintf(`invalid value "%v" for env variable %v : float parsing error`, val, key))
		}
		return v
	}
	return value
}

func outputErr(err string) {
	fmt.Fprintf(flag.CommandLine.Output(), fmt.Sprintf("%v\n", err))
	os.Exit(2)
}
