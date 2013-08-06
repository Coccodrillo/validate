/*
  Copyright (c) 2013 José Carlos Nieto, https://menteslibres.net/xiam

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
	The gosexy/validate package applies a set of validation rules on user-provided
	data.
*/
package validate

import (
	"errors"
	"regexp"
)

var (
	// Common validation rules
	RuleEmail        = regexp.MustCompile(`^[a-zA-Z0-9\+\-\.]+@[a-zA-Z0-9\.\-]+$`)
	RuleURL          = regexp.MustCompile(`^[a-zA-Z0-9]+:\/\/.+`)
	RuleFloat        = regexp.MustCompile(`^[0-9\.]+$`)
	RuleInteger      = regexp.MustCompile(`^[0-9]+$`)
	RuleAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	RuleAlphabetic   = regexp.MustCompile(`^[a-zA-Z]+$`)

	// Default validation errors.
	ErrNotEmail        = errors.New(`Expecting an e-mail.`)
	ErrNotURL          = errors.New(`Expecting an URL.`)
	ErrNotFloat        = errors.New(`Expecting a floating point number.`)
	ErrNotInteger      = errors.New(`Expecting an integer.`)
	ErrNotAlphanumeric = errors.New(`Expecting an alphanumeric string.`)
	ErrNotAlphabetic   = errors.New(`Expecting an alphabetic string.`)
	ErrIsEmpty         = errors.New(`Expecting a non empty value.`)
	ErrEmpty           = errors.New(`Expecting an empty value.`)
)

// Returns error if the provided input is not empty, nil otherwise.
func Empty(input string) error {
	if input != "" {
		return ErrEmpty
	}
	return nil
}

// Returns error if the provided input is empty, nil otherwise.
func NotEmpty(input string) error {
	if input == "" {
		return ErrIsEmpty
	}
	return nil
}

// Returns error if the provided input is not an URL, nil otherwise.
func URL(input string) error {
	if RuleURL.MatchString(input) == false {
		return ErrNotURL
	}
	return nil
}

// Returns error if the provided input is not a floating point number, nil
// otherwise.
func Float(input string) error {
	if RuleFloat.MatchString(input) == false {
		return ErrNotFloat
	}
	return nil
}

// Returns error if the provided input is not an alphanumeric (a-zA-Z0-9)
// string, nil otherwise.
func Alphanumeric(input string) error {
	if RuleAlphanumeric.MatchString(input) == false {
		return ErrNotAlphanumeric
	}
	return nil
}

// Returns error if the provided input is not an alphabetic (a-zA-Z0-9) string,
// nil otherwise.
func Alphabetic(input string) error {
	if RuleAlphabetic.MatchString(input) == false {
		return ErrNotAlphabetic
	}
	return nil
}

// Returns error if the provided input is not an integer value, nil otherwise.
func Integer(input string) error {
	if RuleInteger.MatchString(input) == false {
		return ErrNotInteger
	}
	return nil
}

// Returns error if the provided input is not an email, nil otherwise.
func Email(input string) error {
	if RuleEmail.MatchString(input) == false {
		return ErrNotEmail
	}
	return nil
}

// This function takes an input an applies the given set of validation functions
// in order, each function is a link of the chain. If any validation fails,
// validate.Chain stops and returns the error.
//
// Example:
//
// err := validate.Chain(val, validate.NotEmpty, validate.Email)
//
func Chain(input string, links ...func(string) error) error {
	var err error
	for _, link := range links {
		err = link(input)
		if err != nil {
			return err
		}
	}
	return nil
}

// This function accepts a list of error values (from values or functions) and
// returns the first error found, if any.
//
// Example:
//
// err := validate.Each(
//   validate.Email(userEmail),
//	 validate.Chain(userName, validate.NotEmpty),
// )
func Each(tests ...error) error {
	for i, _ := range tests {
		if tests[i] != nil {
			return tests[i]
		}
	}
	return nil
}

// This function accepts a list of error values (from values or functions) and
// returns an array of errors values, useful for validating all user inputs at
// once.
//
// Example:
//
// err := validate.All(
//   validate.Email(userEmail),
//	 validate.Chain(userName, validate.NotEmpty),
// )
func All(tests ...error) []error {
	res := make([]error, 0, len(tests))

	for i, _ := range tests {
		if tests[i] != nil {
			res = append(res, tests[i])
		}
	}

	return res
}

// This function accepts a list of error values (from values or functions) and
// returns nil if any of the rules is valid.
//
// Example:
//
// err := validate.All(
//   validate.Empty(userPhone),
//   validate.Integer(userPhone),
// )
func Any(tests ...error) error {
	var last error

	for i, _ := range tests {
		if tests[i] == nil {
			return nil
		} else {
			last = tests[i]
		}
	}

	return last
}
