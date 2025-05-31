package main

import (
	"fmt"
	"net/http"
)

func PerseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		processLargePayload()
		fmt.Fprintln(w, "aaa")
		printMemUsage()
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// -------------------- heavy part --------------------
func processLargePayload() string {
	// 50,000 × 1 KiB のバッファ確保
	bufs := make([][]byte, 0, chunkCount)
	for i := 0; i < chunkCount; i++ {
		bufs = append(bufs, make([]byte, chunkSize))
	}
	println("perse")

	// 解析フェーズ（ダミー）
	parse(bufs)

	return "done"
}

func parse2(_ [][]byte) {
	// 実アプリでは CSV パースなどを行う
}

const (
	chunkCount = 50_000
	chunkSize  = 1_024
)
