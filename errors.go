package python

import "fmt"

type RunError struct {
	Code int
	File string
}

func (e RunError) Error() string {
	return fmt.Sprintf("python: error %d executing script %s", e.Code, e.File)
}

func errCode(ret int) error {
	if ret == 0 {
		return nil
	}
	return Error{Code: ret}
}

type Error struct {
	Code int
}

func (e Error) Error() string {
	return fmt.Sprintf("python: C-Python error code %d", e.Code)
}
