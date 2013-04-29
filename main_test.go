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
	"testing"
)

func TestValidateEmail(t *testing.T) {
	var err error

	err = Email("user@example.com")
	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = Email("use`r@example.com")
	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateURL(t *testing.T) {
	var err error

	err = URL("ftp://foo:bar@archive.org/")

	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = URL("user@example.com")

	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateFloat(t *testing.T) {
	var err error

	err = Float("1.23")

	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = Float("One")

	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateInteger(t *testing.T) {
	var err error

	err = Integer("43")

	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = Integer("1.23")

	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateAlphanumeric(t *testing.T) {
	var err error

	err = Alphanumeric("ABc123")

	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = Alphanumeric("ABc$23")

	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateAlphabetic(t *testing.T) {
	var err error

	err = Alphabetic("ABc")

	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	err = Alphabetic("4ba")

	if err == nil {
		t.Fatalf("Failed validation")
	}
}

func TestValidateChain(t *testing.T) {
	var err error

	fnOk := func(s string) error {
		return nil
	}

	err = Chain("user@example.com", Email, fnOk)
	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}

	fnErr := func(s string) error {
		return nil
	}
	err = Chain("user@example.com", NotEmpty, Email, fnErr, fnOk)
	if err != nil {
		t.Fatalf("Failed validation: %s", err.Error())
	}
}

func TestValidateEach(t *testing.T) {
	var err error

	err = Each(
		Email("user@example.com"),
		Email("userexample.com"),
		Email("user@example.com"),
	)

	if err == nil {
		t.Fatalf("Expecting error.")
	}

	err = Each(
		Email("user@example.com"),
		NotEmpty("hola"),
		Float("1.23"),
	)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestValidateAll(t *testing.T) {
	var err []error

	err = All(
		Email("user@example.com"),
		Email("userexample.com"),
		Email("user@example.com"),
	)

	if len(err) == 0 {
		t.Fatalf("Expecting error.")
	}

	err = All(
		Email("user@example.com"),
		NotEmpty("hola"),
		Float("1.23"),
	)

	if len(err) != 0 {
		t.Fatalf(err[0].Error())
	}
}

func TestValidateAny(t *testing.T) {
	var err error

	err = Any(
		Email("user@example.com"),
		Email("userexample.com"),
		Email("user@example.com"),
	)

	if err != nil {
		t.Fatalf("Expecting at least one valid e-mail.")
	}

	err = Any(
		Email("user@example.com"),
		NotEmpty("hola"),
		Float("1.23"),
	)

	if err != nil {
		t.Fatalf(err.Error())
	}

	err = Any(
		Email("userexample.com"),
		Email("123"),
		Float("a"),
	)

	if err == nil {
		t.Fatalf("Expecting an error")
	}
}
