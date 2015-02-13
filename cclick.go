package explode

/*
#cgo pkg-config: click-0.4
#cgo pkg-config: glib-2.0 gobject-2.0

#include <click-0.4/click.h>
*/
import "C"

import (
	"fmt"
	"runtime"
)

type CClickUser struct {
	cref *C.ClickUser
}

func (ccu *CClickUser) CInit(holder interface{}) error {
	var gerr *C.GError
	cref := C.click_user_new_for_user(nil, nil, &gerr)
	defer C.g_clear_error(&gerr)
	if gerr != nil {
		return fmt.Errorf("faild to make ClickUser: %s", C.GoString((*C.char)(gerr.message)))
	}
	ccu.cref = cref
	runtime.SetFinalizer(holder, func(interface{}) {
        // ccu.cref = nil
		C.g_object_unref((C.gpointer)(cref))
	})
	return nil
}
