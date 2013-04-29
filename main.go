/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

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

package validate

import (
	"errors"
	"regexp"
)

// Don't be too restrictive, unless you really need it.
var (
	RuleEmail        = regexp.MustCompile(`^[a-zA-Z0-9\+\-\.]+@[a-zA-Z0-9\.\-]+$`)
	RuleURL          = regexp.MustCompile(`^[a-zA-Z0-9]+:\/\/.+`)
	RuleFloat        = regexp.MustCompile(`^[0-9\.]+$`)
	RuleInteger      = regexp.MustCompile(`^[0-9]+$`)
	RuleAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	RuleAlphabetic   = regexp.MustCompile(`^[a-zA-Z]+$`)

	ErrNotEmail        = errors.New(`Expecting an e-mail.`)
	ErrNotURL          = errors.New(`Expecting an URL.`)
	ErrNotFloat        = errors.New(`Expecting a floating point number (0-9 and point).`)
	ErrNotInteger      = errors.New(`Expecting an integer number.`)
	ErrNotAlphanumeric = errors.New(`Expecting alphanumeric.`)
	ErrNotAlphabetic   = errors.New(`Expecting an alphabetic string.`)

	ErrIsEmpty = errors.New(`Expecting a non empty value.`)
	ErrEmpty   = errors.New(`Expecting an empty value.`)
)

func Empty(input string) error {
	if input != "" {
		return ErrEmpty
	}
	return nil
}

func NotEmpty(input string) error {
	if input == "" {
		return ErrIsEmpty
	}
	return nil
}

func URL(input string) error {
	if RuleURL.MatchString(input) == false {
		return ErrNotURL
	}
	return nil
}

func Float(input string) error {
	if RuleFloat.MatchString(input) == false {
		return ErrNotFloat
	}
	return nil
}

func Alphanumeric(input string) error {
	if RuleAlphanumeric.MatchString(input) == false {
		return ErrNotAlphanumeric
	}
	return nil
}

func Alphabetic(input string) error {
	if RuleAlphabetic.MatchString(input) == false {
		return ErrNotAlphabetic
	}
	return nil
}

func Integer(input string) error {
	if RuleInteger.MatchString(input) == false {
		return ErrNotInteger
	}
	return nil
}

func Email(input string) error {
	if RuleEmail.MatchString(input) == false {
		return ErrNotEmail
	}
	return nil
}

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

func Each(tests ...error) error {
	for i, _ := range tests {
		if tests[i] != nil {
			return tests[i]
		}
	}
	return nil
}

func All(tests ...error) []error {
	res := make([]error, 0, len(tests))

	for i, _ := range tests {
		if tests[i] != nil {
			res = append(res, tests[i])
		}
	}

	return res
}

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
