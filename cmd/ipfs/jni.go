package main

// #include <stddef.h>
// #include <stdint.h>
// #include <stdlib.h>
import "C"
import (
	"sync"

	"github.com/ClarkGuan/jni"
	"github.com/ipfs/go-ipfs/cmd/ipfs/msgs"
)

var (
	mtx sync.Mutex
)

func getBytes(env uintptr, iarr uintptr) []byte {
	var buf []byte
	arrLen := jni.Env(env).GetArrayLength(iarr)
	for i := 0; i < arrLen; i++ {
		b := jni.Env(env).GetByteArrayElement(iarr, i)
		buf = append(buf, b)
	}
	return buf
}

//export Java_io_zzv_jni_Main_call
func Java_io_zzv_jni_Main_call(env uintptr, clazz uintptr, iarr uintptr) uintptr {
	mtx.Lock()
	defer mtx.Unlock()

	msgs.SetEnv(env, clazz)
	rarr := call(getBytes(env, iarr))
	jarr := jni.Env(env).NewByteArray(len(rarr))
	jni.Env(env).SetByteArrayRegion(jarr, 0, rarr)
	return jarr
}
