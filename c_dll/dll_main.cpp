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