package errors

import "bytes"

type bucket struct {
	cause    error   // error cause
	text     string  // error text (msg)
	code     *uint32 // error code
	location *string // error location (file:line)
	external bool    // error is external text
}

// Error method implement error interface return full message a error
func (b *bucket) Error() string {
	if b.external {
		return b.text
	}

	buf := new(bytes.Buffer)
	buf.WriteString(b.text)

	err := Unwrap(b)

	for err != nil {
		buf.WriteString(delimiter)

		if e, ok := err.(interface{ string() string }); ok {
			buf.WriteString(e.string())
		} else {
			buf.WriteString(err.Error())
		}

		err = Unwrap(err)
	}

	return buf.String()
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (b *bucket) Unwrap() error {
	return b.cause
}

func (b *bucket) string() string {
	return b.text
}

func (b *bucket) getCode() *uint32 {
	return b.code
}

func (b *bucket) setCode(code *uint32) error {
	b.location = getLocation(3)
	b.code = code

	return b
}

func (b *bucket) setLocation() error {
	b.location = getLocation(3)

	return b
}

func (b *bucket) getLocation() string {
	if b.location == nil {
		return ""
	}

	return *b.location
}
