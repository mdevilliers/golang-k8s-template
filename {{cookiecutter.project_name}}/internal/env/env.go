package env

import (
	"os"
	"strconv"
)

// LookUpWithDefaultStr return the found value or the default
// This really should be in a shared library!
func LookUpWithDefaultStr(name string, defaultValue string) string {

	value, ok := os.LookupEnv(name)

	if ok {
		return value
	}

	return defaultValue
}

// LookUpWithDefaultBool return the found value or the default
// This really should be in a shared library!
func LookUpWithDefaultBool(name string, defaultValue bool) bool {
	value, ok := os.LookupEnv(name)

	if ok {

		b, err := strconv.ParseBool(value)

		if err != nil {
			return defaultValue
		}
		return b
	}

	return defaultValue

}
