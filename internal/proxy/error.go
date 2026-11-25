package proxy

import (
	"fmt"
)

var ErrServerError = fmt.Errorf("Internal Server Error")
var ErrServerIsUnreachable = fmt.Errorf("Server is unreachable")
