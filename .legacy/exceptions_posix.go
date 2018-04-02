// +build !windows

package python

// #include "go-python.h"
import "C"

// int PySignal_SetWakeupFd(int fd)
// This utility function specifies a file descriptor to which a '\0' byte will be written whenever a signal is received. It returns the previous such file descriptor. The value -1 disables the feature; this is the initial state. This is equivalent to signal.set_wakeup_fd() in Python, but without any error checking. fd should be a valid file descriptor. The function should only be called from the main thread.
func PySignal_SetWakeupFd(fd int) int {
	return int(C.PySignal_SetWakeupFd(C.int(fd)))
}
