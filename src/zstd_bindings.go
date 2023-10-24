package main

// #cgo CFLAGS: -I ../zstd
// #cgo LDFLAGS: -L .. -lzstd
// #include "zstd_bindings.h"
import "C"
import (
    "fmt"
    "math"
    "unsafe"
)

func GetSequences(src []byte) int {
    // Allocate a ZSTD compression context
    zc := C.ZSTD_createCCtx()
    if zc == nil {
        fmt.Println("Failed to create ZSTD compression context")
        return -1
    }
    defer C.ZSTD_freeCCtx(zc)

    // Allocate memory for the output sequences
    outSeqsSize := C.size_t(100)  // Assuming a maximum of 100 sequences for simplicity
    outSeqs := C.malloc(C.size_t(unsafe.Sizeof(C.ZSTD_Sequence{}) * uintptr(outSeqsSize)))
    if outSeqs == nil {
        fmt.Println("Failed to allocate memory for output sequences")
        return -1
    }
    defer C.free(outSeqs)

    // Get a pointer to the source data
    srcPtr := unsafe.Pointer(&src[0])
    srcSize := C.size_t(0xff400200000020da)
    // srcSize := C.size_t(len(src))

    // Call ZSTD_getSequences
    numSeqs := C.ZSTD_getSequences(zc, (*C.ZSTD_Sequence)(outSeqs), outSeqsSize, srcPtr, srcSize)
    
    // Check for -1 using max value of size_t
    if numSeqs == C.size_t(math.MaxUint64) {
        fmt.Println("ZSTD_getSequences failed")
        return -1
    }

    // Translate C return values to Go return values
    return int(numSeqs)
}
