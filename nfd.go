package nfd_go

/*
#cgo CFLAGS: -Ithird_party/nativefiledialog/src/include

#cgo linux openbsd freebsd CXXFLAGS: -DLINUX
#cgo linux openbsd freebsd pkg-config: --cflags gtk+-3.0 glib-2.0 --libs glib-2.0

#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework AppKit

#cgo windows CXXFLAGS: -DWINDOWS
#cgo windows LDFLAGS: -lole32 -lshell32 -luuid

#include <stdlib.h>
#include "nfd.h"
#include "third_party/nativefiledialog/src/nfd_common.c"

#if defined(DARWIN)
#include "third_party/nativefiledialog/src/nfd_cocoa.m"
#endif //DARWIN

#if defined(WINDOWS)
#include "third_party/nativefiledialog/src/nfd_win.cpp"
#endif //WINDOWS

#if defined(LINUX)
#include "third_party/nativefiledialog/src/nfd_gtk.c"
#endif //LINUX

*/
import "C"
import (
	"errors"
	"unsafe"
)

func PickFolder(defaultPath string) (res string, err error)  {
	cDefaultPath := cString(defaultPath)

	var cOutPath *C.char
	switch C.NFD_PickFolder(cDefaultPath, &cOutPath) {
	case C.NFD_OKAY:
		res = C.GoString(cOutPath)
		free(cOutPath)
	case C.NFD_CANCEL:
		// empty
	default:
		err = getError()
	}

	free(cDefaultPath)
	return
}

func cString(str string) *C.char {
	if str == "" {
		return nil
	}
	return C.CString(str)
}

func free(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func getError() error {
	return errors.New(C.GoString(C.NFD_GetError()))
}