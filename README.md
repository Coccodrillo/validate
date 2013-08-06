# gosexy/validation

The `gosexy/validation` package applies a set of validation rules on
user-provided data, this data could be from any source, including web forms or
web-based API parameters.

## Installation

```sh
go get -u menteslibres.net/gosexy/validate
```

## Usage

Import the validation package.

```go
import "menteslibres.net/gosexy/validate"
```

Use validation functions to make sure the user is providing a valid e-mail
address. See the full list of validation rules in the [docs][1].

```go
userEmail := "user@example.com"
err := validate.Email(userEmail)
if err != nil {
  log.Fatalf("This is not an e-mail!")
}
```

Note that many validations can be applied at once on the same user input using
`validate.Chain()`.

```go
err := validate.Chain(userEmail, validate.NotEmpty, validate.Email)
if err != nil {
  log.Fatalf("This is not an e-mail!")
}
log.Printf("Got an e-mail address.")
```

Any function of the form `func(string) error` can be used along with
`validate.Chain()`, so you can add custom validations.

```go
fnCustomValidation = function(s string) error {
  // Some validation
  return nil
}
err := validate.Chain(userEmail, validate.NotEmpty, validate.Email, fnCustomValidation)
if err != nil {
  log.Fatalf("This is not an e-mail!")
}
log.Printf("Got an e-mail address.")
```

If you need to check many rules use `validate.Each()` and pass some `error`
values, the first `error` value found will be returned.

```go
firstErr := validate.Each(
  validate.Email(userEmail),
  validate.NotEmpty(userName),
)
if firstErr == nil {
  // Data is clean!
}
```

The function `validate.All()` makes sure all validations are tested, it returns
an array of errors (`[]error{}`) with the failed tests. Useful for checking
all user inputs at once.

```go
errs := validate.All(
  validate.Email(userEmail),
  validate.NotEmpty(userName),
)
```

Another useful function is `validate.Any()`, will return `nil` if any of the
rules return `nil`.

```go
err := validate.Any(userAge, validate.Empty, validate.Integer)
if err == nil {
  // Value is empty or an e-mail.
}
```

## Documentation

See the [online docs][1].

## License

>  Copyright (c) 2013 José Carlos Nieto, https://menteslibres.net/xiam
>
>  Permission is hereby granted, free of charge, to any person obtaining
>  a copy of this software and associated documentation files (the
>  "Software"), to deal in the Software without restriction, including
>  without limitation the rights to use, copy, modify, merge, publish,
>  distribute, sublicense, and/or sell copies of the Software, and to
>  permit persons to whom the Software is furnished to do so, subject to
>  the following conditions:
>
>  The above copyright notice and this permission notice shall be
>  included in all copies or substantial portions of the Software.
>
>  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
>  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
>  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
>  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
>  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
>  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
>  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[1]: http://godoc.org/menteslibres.net/gosexy/validate
