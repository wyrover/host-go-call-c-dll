package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <windows.h>


// go callback 函数
int in_go_callback(char* a, char*b);

typedef int (*in_c_function)(char* a, char* b);
// C callback 函数指针
in_c_function fn_in_c_function;


static inline void Hello(char* str) {
    printf("%s", str);
}

static inline int register_functions()
{
    int retval = -1;
    typedef int (*fncdll)();
    HINSTANCE hDLL;

    do 
    {
        
        hDLL = LoadLibraryW(L".\\c_dll.dll");
        if (hDLL == NULL) {
            break;
        }

        fncdll fn = (fncdll)GetProcAddress(hDLL, "register_functions");
        if (fn == NULL) {
            break;
        }

        retval = fn(in_go_callback, &fn_in_c_function);      

    } while (0);
   
   
    //if (hDLL) {
    //    FreeLibrary(hDLL);       
    //} 

    return retval;
}

static inline int call_c_function(char* a, char* b) {
    return fn_in_c_function(a, b);
}






*/
import "C"

import (
    "unsafe"
    "fmt"
)

func main() {
    
    if C.register_functions() != 0 {
        panic("Load dll failed")
    }

    aa := C.CString("go call c arg1")
    bb := C.CString("go call c arg2")

    defer C.free(unsafe.Pointer(aa))
    defer C.free(unsafe.Pointer(bb))

    
    for i := 0; i < 10; i++ {
        C.call_c_function(aa, bb)
    }
}

//export in_go_callback
func in_go_callback(a *C.char, b *C.char) C.int {
    fmt.Println(C.GoString(a))
    fmt.Println(C.GoString(b))
    return 0
}