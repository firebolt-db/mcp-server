package args

import (
	"errors"
	"fmt"
)

// String retrieves a required string argument from the provided map.
func String(arguments map[string]any, name string) (string, error) {

	val, ok := arguments[name]
	if !ok || val == nil {
		return "", fmt.Errorf("required argument %s not provided", name)
	}

	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("argument %s is not a string", name)
	}

	return str, nil
}

// MaybeString retrieves an optional string argument from the provided map.
func MaybeString(arguments map[string]any, name string) (*string, error) {

	val, ok := arguments[name]
	if !ok || val == nil {
		return nil, nil
	}

	str, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("argument %s is not a string", name)
	}

	return &str, nil
}

// Strings retrieves multiple required string arguments from the provided map.
func Strings(arguments map[string]any, names ...string) ([]string, error) {

	values := make([]string, len(names))
	errs := make([]error, len(names))

	for i, name := range names {
		values[i], errs[i] = String(arguments, name)
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return values, nil
}

// MaybeStrings retrieves multiple optional string arguments from the provided map.
func MaybeStrings(arguments map[string]any, names ...string) ([]*string, error) {

	values := make([]*string, len(names))
	errs := make([]error, len(names))

	for i, name := range names {
		values[i], errs[i] = MaybeString(arguments, name)
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return values, nil
}
