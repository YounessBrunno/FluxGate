package proxy

import "errors"

var ErrServerError = errors.New("internal server error")
var ErrServerIsUnreachable = errors.New("server is unreachable")