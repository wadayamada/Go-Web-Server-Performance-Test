package main

import (
	"fmt"
	"runtime"
)

func printMemUsage() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Process (Go runtime)\n")
	fmt.Printf("  Alloc     : %.2f MB\n", float64(ms.Alloc)/1024/1024)
	fmt.Printf("  Sys       : %.2f MB\n", float64(ms.Sys)/1024/1024)
	fmt.Printf("  HeapAlloc : %.2f MB\n", float64(ms.HeapAlloc)/1024/1024)
	fmt.Printf("  HeapSys   : %.2f MB\n", float64(ms.HeapSys)/1024/1024)
	fmt.Printf("  NextGC    : %.2f MB\n", float64(ms.NextGC)/1024/1024)
	fmt.Printf("  NumGC     : %d\n", ms.NumGC)
}

const (
	runs     = 1000  // 計測回数
	chunkCnt = 50000 // 生成バッファ数
	chunkSz  = 1024  // 1 バッファのサイズ
)
