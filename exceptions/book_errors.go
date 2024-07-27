package exceptions

import "errors"

var (
	ErrBookIdDup = errors.New("book duplicate id")
)