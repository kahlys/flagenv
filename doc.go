// Package flagenv is an extension of https://golang.org/pkg/flag/ adding environment variables.
//
// Usage
//
// Define flags using flag.String(), Bool(), Int(), etc.
//
// This declares a string flag -option or from OPTION env variable, stored in the pointer
// optFlag, with type *string:
// 	import "flagenv"
// 	var optFlag = flagenv.String("option", "OPTION", "value", "help message")
//
// After all flags are defined, call
// 	flagenv.Parse()
// to parse the command line into the defined flags.
package flagenv
