package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

// From Form creates a custom form struct, embeds a url.Values object
type From struct {
	url.Values
	Errors errors
}

// Valid return true if there are no errors, otherwise false
func (f *From) Valid() bool {
	return len(f.Errors) == 0
}

// New initialized a form struct
func New(data url.Values) *From {
	return &From{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *From) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *From) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// MinLength check the minimum length
func (f *From) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long. %d given", length, len(x)))
		return false
	}
	return true
}

func (f *From) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
