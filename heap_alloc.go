package main

import (
	"fmt"
	"time"
)

// -------------------- heavy part --------------------
func use_heap() string {
	// 50,000 × 1 KiB のバッファ確保
	bufs := make([][]byte, 0, chunkCnt)
	for i := 0; i < chunkCnt; i++ {
		bufs = append(bufs, make([]byte, chunkSz))
	}
	println("perse")

	return "done"
}

/*----------------- benchmark ------------------*/
func main() {
	var total float64
	// レスポンス速度を入れる配列
	resps := make([]float64, 0, runs)

	for i := 1; i <= runs; i++ {
		start := time.Now()
		res := use_heap()
		printMemUsage()
		sec := time.Since(start).Seconds()
		total += sec
		resps = append(resps, sec)
		fmt.Printf("Run %3d: %.6f seconds (%s)\n", i, sec, res)
	}
	fmt.Println("Response times:", resps)
	fmt.Printf("Average over %d runs: %.6f seconds\n",
		runs, total/float64(runs))
}
