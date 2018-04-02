// +build !windows

package python

//#include <stdlib.h>
//#include <string.h>
//#include <stdio.h>
import "C"

import (
	"os"
	"unsafe"
)

// file2py opens a stdC file from a Go os.File.  Note the returned file has
// been newly opened: the caller must close it with C.fclose(retval).
func file2py(f *os.File, mode string) *C.FILE {
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	fd := f.Fd()
	file := C.fdopen(C.int(fd), cmode)
	return file
}
