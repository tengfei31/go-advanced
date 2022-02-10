package cgo

import "C"

//export number_add_mod
func number_add_mod(a C.int, b C.int, mod C.int) C.int {
	return (a + b) % mod
}
