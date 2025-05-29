package trn

/*
#cgo CFLAGS: -g -Wall -I/home/rds/include
#cgo LDFLAGS: -L/home/rds/lib -lrds_trn
#include <stdio.h>
#include <stdlib.h>
#include <rds_trn.h>

  void trn_Register(char *name) {
    trn_register(name) ;
  }

  void trace(char *input) {
     Trace("%s",input) ;
  }
  void inform(char *input) {
     Inform("%s",input) ;
  }
  void alert(char *input) {
     Alert("%s",input) ;
  }
*/
import "C"
import "fmt"
import "unsafe"

func Register(name string) {
	cName := C.CString(name)
	C.trn_Register(cName)
	C.free(unsafe.Pointer(cName))
}

func Trace(format string, args ...interface{}) {
	msg := C.CString(fmt.Sprintf(format, args...))
	C.trace(msg)
	C.free(unsafe.Pointer(msg))
}
func Alert(format string, args ...interface{}) {
	msg := C.CString(fmt.Sprintf(format, args...))
	C.alert(msg)
	C.free(unsafe.Pointer(msg))
}
func Inform(format string, args ...interface{}) {
	msg := C.CString(fmt.Sprintf(format, args...))
	C.inform(msg)
	C.free(unsafe.Pointer(msg))
}
