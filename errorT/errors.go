package errort

import "errors"

var ErrNotFound = errors.New("not found")
var ErrOutOfRange = errors.New("out of range")
var ErrInvalidParam = errors.New("invalid param")
var ErrInvalidParamType = errors.New("invalid param type")
var ErrInvalidParamValue = errors.New("invalid param value")
var ErrInvalidParamLength = errors.New("invalid param length")
var ErrInvalidParamFormat = errors.New("invalid param format")
var ErrZeroValue = errors.New("zero value")
