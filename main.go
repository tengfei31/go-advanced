package main

/*
#include <stdint.h>

int64_t myadd(int64_t a, int64_t b) {
    return a+b;
}
*/
//import "C"

import (
	"go-advanced/assembly"
	"sync"
)

func main() {
	//fmt.Println(assembly.AsmCallCAdd(uintptr(unsafe.Pointer(C.myadd)), 123, 456))
	var waitGo sync.WaitGroup
	waitGo.Add(1)
	go func(wg *sync.WaitGroup) {
		assembly.Test()
		wg.Done()
	}(&waitGo)
	waitGo.Wait()
}
