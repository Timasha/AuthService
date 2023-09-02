package errs

import "fmt"

var ErrWrongSingingMethod error = fmt.Errorf("wrong singing method")
