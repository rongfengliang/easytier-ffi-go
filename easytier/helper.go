package easytier

import "unsafe"

func CString(s string) ([]byte, *byte) {
	b := append([]byte(s), 0)
	return b, &b[0]
}

func CStringPtr(s string) uintptr {
	b := append([]byte(s), 0)
	return uintptr(unsafe.Pointer(&b[0]))
}

func CStrToGoStr(cstr *byte) string {
	return unsafe.String(cstr, strlen(cstr))
}

func strlen(p *byte) int {
	n := 0
	for *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(n))) != 0 {
		n++
	}
	return n
}
