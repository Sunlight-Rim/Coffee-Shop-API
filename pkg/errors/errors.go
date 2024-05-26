// Package errors implements functions for handling errors.
package errors

import (
	"fmt"
	"runtime"
)

var delimiter = ": "

// SetDelimiter replaces the default delimiter ": " with a custom one
//
// Warning, this operation is not safe for concurrent use, use in init function
func SetDelimiter(sep string) {
	delimiter = sep
}

// New create an error based on the error text and an optional error code parameter
func New(text string, code ...uint32) error {
	return &bucket{
		external: false,
		text:     text,
		location: getLocation(2),
		code:     getFirstCode(code),
		cause:    nil,
	}
}

// String returns the error text, the method differs from Error() in that it will return the original error text if she was created by this package
func String(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	if e, ok := err.(interface{ string() string }); ok {
		return e.string(), true
	}

	return err.Error(), false
}

// SetCode sets the error uint32 code for the given error
func SetCode(err error, code uint32) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(interface{ setCode(code *uint32) error }); ok {
		return e.setCode(&code)
	}

	return &bucket{
		external: true,
		text:     err.Error(),
		location: getLocation(2),
		code:     &code,
		cause:    err,
	}
}

// GetCode will return an error code if it was set
func GetCode(err error) (uint32, bool) {
	if err == nil {
		return 0, false
	}

	if e, ok := err.(interface{ getCode() *uint32 }); ok {
		if code := e.getCode(); code != nil {
			return *code, true
		}

		return 0, false
	}

	return 0, false
}

// SetLocation sets path file and code line where was called
func SetLocation(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(interface{ setLocation() error }); ok {
		return e.setLocation()
	}

	return &bucket{
		external: true,
		text:     err.Error(),
		location: getLocation(2),
		code:     nil,
		cause:    err,
	}
}

// GetLocation return a error path file and code line if was set
func GetLocation(err error) (string, bool) {
	if e, ok := err.(interface{ getLocation() string }); ok {
		return e.getLocation(), ok
	}

	return "", false
}

// Errorf create an error with string formatting
func Errorf(format string, args ...any) error {
	return &bucket{
		external: false,
		text:     fmt.Sprintf(format, args...),
		location: getLocation(2),
		code:     nil,
		cause:    nil,
	}
}

// Wrap creates a new error and adds the original one to the context as a cause
func Wrap(err error, text string) error {
	return wrap(err, text)
}

// Wrap creates a new error and adds the original one to the context as a cause with string formatting
func Wrapf(err error, format string, args ...any) error {
	return wrap(err, fmt.Sprintf(format, args...))
}

func wrap(err error, text string) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(interface{ getCode() *uint32 }); ok {
		return &bucket{
			external: false,
			text:     text,
			location: getLocation(3),
			code:     e.getCode(),
			cause:    err,
		}
	}

	return &bucket{
		external: false,
		text:     text,
		location: getLocation(3),
		code:     nil,
		cause:    err,
	}
}

func getFirstCode(code []uint32) *uint32 {
	var c *uint32

	if len(code) != 0 {
		c = &code[0]
	}

	return c
}

func getLocation(skip int) *string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		location := fmt.Sprintf("%s:%d", file, line)

		return &location
	}

	return nil
}
