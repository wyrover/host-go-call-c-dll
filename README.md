# host-go-call-c-dll

go 程序作为宿主加载 C++ DLL 双向通信。

## GO

```go
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
```

## DLL 

```cpp
#ifdef DLL_EXPORTS
#define DLL_API __declspec(dllexport)
#else
#define DLL_API __declspec(dllimport)
#endif

#include <stdlib.h>
#include <stdio.h>

#ifdef __cplusplus
extern "C" {
#endif


typedef int (*in_go_function)(char* a, char* b);
in_go_function fn_in_go_function;

typedef int (*in_c_function)(char* a, char* b);


int in_c_callback(char* a, char* b) {
    printf("%s\n", a);
    printf("%s\n", b);
}


DLL_API int register_functions(in_go_function fn_go, in_c_function* fn_c)
{
    fn_in_go_function = fn_go;
    *fn_c = in_c_callback;
    
    for (int i = 0; i < 10; ++i) {    
        fn_in_go_function("c call go arg1", "c call go arg2");
    }
    
    return 0;
}


#ifdef __cplusplus
}
#endif
```