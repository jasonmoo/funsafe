package funsafe

import (
	"reflect"
	"unsafe"
)

// Unset read-only flags on `reflect.Value`
//
func MakeSettable(v *reflect.Value) {

	// https://golang.org/src/reflect/value.go#L27
	// memory layout of members of reflect.Value struct
	type value struct {
		typ  uintptr
		ptr  uintptr
		flag uintptr
	}

	// https://golang.org/src/reflect/value.go#L71
	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	const roFlag uintptr = 1<<5 | 1<<6

	(*value)(unsafe.Pointer(v)).flag &= ^roFlag

}
