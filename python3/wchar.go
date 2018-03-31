/*
gowchar library.
https://github.com/orofarne/gowchar/

Copyright (c) 2013, Maxim Dementyev. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of the author. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package python3

/*
#include <wchar.h>

const size_t SIZEOF_WCHAR_T = sizeof(wchar_t);

void gowchar_set (wchar_t *arr, int pos, wchar_t val) {
	arr[pos] = val;
}

wchar_t gowchar_get (wchar_t *arr, int pos) {
	return arr[pos];
}
*/
import "C"

import (
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

var sizeof_WCHAR_T C.size_t = C.size_t(C.SIZEOF_WCHAR_T)

func stringToWcharT(s string) (*C.wchar_t, C.size_t) {
	switch sizeof_WCHAR_T {
	case 2:
		return stringToWchar2(s) // Windows
	case 4:
		return stringToWchar4(s) // Unix
	default:
		panic(fmt.Sprintf("Invalid sizeof(wchar_t) = %v", sizeof_WCHAR_T))
	}
	panic("?!!")
}

func wcharTToString(s *C.wchar_t) (string, error) {
	switch sizeof_WCHAR_T {
	case 2:
		return wchar2ToString(s) // Windows
	case 4:
		return wchar4ToString(s) // Unix
	default:
		panic(fmt.Sprintf("Invalid sizeof(wchar_t) = %v", sizeof_WCHAR_T))
	}
	panic("?!!")
}

func wcharTNToString(s *C.wchar_t, size C.size_t) (string, error) {
	switch sizeof_WCHAR_T {
	case 2:
		return wchar2NToString(s, size) // Windows
	case 4:
		return wchar4NToString(s, size) // Unix
	default:
		panic(fmt.Sprintf("Invalid sizeof(wchar_t) = %v", sizeof_WCHAR_T))
	}
	panic("?!!")
}

// Windows
func stringToWchar2(s string) (*C.wchar_t, C.size_t) {
	var slen int
	s1 := s
	for len(s1) > 0 {
		r, size := utf8.DecodeRuneInString(s1)
		if er, _ := utf16.EncodeRune(r); er == '\uFFFD' {
			slen += 1
		} else {
			slen += 2
		}
		s1 = s1[size:]
	}
	slen++ // \0
	res := C.malloc(C.size_t(slen) * sizeof_WCHAR_T)
	var i int
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r1, r2 := utf16.EncodeRune(r); r1 != '\uFFFD' {
			C.gowchar_set((*C.wchar_t)(res), C.int(i), C.wchar_t(r1))
			i++
			C.gowchar_set((*C.wchar_t)(res), C.int(i), C.wchar_t(r2))
			i++
		} else {
			C.gowchar_set((*C.wchar_t)(res), C.int(i), C.wchar_t(r))
			i++
		}
		s = s[size:]
	}
	C.gowchar_set((*C.wchar_t)(res), C.int(slen-1), C.wchar_t(0)) // \0
	return (*C.wchar_t)(res), C.size_t(slen)
}

// Unix
func stringToWchar4(s string) (*C.wchar_t, C.size_t) {
	slen := utf8.RuneCountInString(s)
	slen++ // \0
	res := C.malloc(C.size_t(slen) * sizeof_WCHAR_T)
	var i int
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		C.gowchar_set((*C.wchar_t)(res), C.int(i), C.wchar_t(r))
		s = s[size:]
		i++
	}
	C.gowchar_set((*C.wchar_t)(res), C.int(slen-1), C.wchar_t(0)) // \0
	return (*C.wchar_t)(res), C.size_t(slen)
}

// Windows
func wchar2ToString(s *C.wchar_t) (string, error) {
	var i int
	var res string
	for {
		ch := C.gowchar_get(s, C.int(i))
		if ch == 0 {
			break
		}
		r := rune(ch)
		i++
		if !utf16.IsSurrogate(r) {
			if !utf8.ValidRune(r) {
				err := fmt.Errorf("Invalid rune at position %v", i)
				return "", err
			}
			res += string(r)
		} else {
			ch2 := C.gowchar_get(s, C.int(i))
			r2 := rune(ch2)
			r12 := utf16.DecodeRune(r, r2)
			if r12 == '\uFFFD' {
				err := fmt.Errorf("Invalid surrogate pair at position %v", i-1)
				return "", err
			}
			res += string(r12)
			i++
		}
	}
	return res, nil
}

// Unix
func wchar4ToString(s *C.wchar_t) (string, error) {
	var i int
	var res string
	for {
		ch := C.gowchar_get(s, C.int(i))
		if ch == 0 {
			break
		}
		r := rune(ch)
		if !utf8.ValidRune(r) {
			err := fmt.Errorf("Invalid rune at position %v", i)
			return "", err
		}
		res += string(r)
		i++
	}
	return res, nil
}

// Windows
func wchar2NToString(s *C.wchar_t, size C.size_t) (string, error) {
	var i int
	var res string
	N := int(size)
	for i < N {
		ch := C.gowchar_get(s, C.int(i))
		if ch == 0 {
			break
		}
		r := rune(ch)
		i++
		if !utf16.IsSurrogate(r) {
			if !utf8.ValidRune(r) {
				err := fmt.Errorf("Invalid rune at position %v", i)
				return "", err
			}

			res += string(r)
		} else {
			if i >= N {
				err := fmt.Errorf("Invalid surrogate pair at position %v", i-1)
				return "", err
			}
			ch2 := C.gowchar_get(s, C.int(i))
			r2 := rune(ch2)
			r12 := utf16.DecodeRune(r, r2)
			if r12 == '\uFFFD' {
				err := fmt.Errorf("Invalid surrogate pair at position %v", i-1)
				return "", err
			}
			res += string(r12)
			i++
		}
	}
	return res, nil
}

// Unix
func wchar4NToString(s *C.wchar_t, size C.size_t) (string, error) {
	var i int
	var res string
	N := int(size)
	for i < N {
		ch := C.gowchar_get(s, C.int(i))
		r := rune(ch)
		if !utf8.ValidRune(r) {
			err := fmt.Errorf("Invalid rune at position %v", i)
			return "", err
		}
		res += string(r)
		i++
	}
	return res, nil
}
