package main

import (
	"fmt"
	"time"
)

//const (
//	runs     = 1000  // 計測回数
//	chunkCnt = 50000 // 生成バッファ数
//	chunkSz  = 1024  // 1 バッファのサイズ
//)

/*----------------- heavy part -----------------*/
//func buildLargeHeapResponse() string {
//	bufs := make([][]byte, 0, chunkCnt)
//	for i := 0; i < chunkCnt; i++ {
//		bufs = append(bufs, make([]byte, chunkSz)) // ヒープ確保
//	}
//	printMemUsage()
//	parse(bufs) // ダミー解析
//	return "done"
//}

// スタック上で Item を作って簡単な計算を行う
//func buildLargeResponse() uint64 {
//	var total uint64
//	for i := 0; i < chunkCnt; i++ {
//		// Item は値として生成→ポインタを外に出さないので escape しない
//		it := createItemAddress()
//		total += it.data[0] // ダミー集計
//	}
//	return total
//}
//
//func createItem() Item {
//	it := Item{}
//	// 適当に 0..127 を書き込む
//	for j := 0; j < len(it.data); j++ {
//		it.data[j] = uint64(j)
//	}
//	return it // スタック上の Item を返す
//}
//
//func createItemAddress() *Item {
//	it := Item{}
//	// 適当に 0..127 を書き込む
//	for j := 0; j < len(it.data); j++ {
//		it.data[j] = uint64(j)
//	}
//	return &it // スタック上の Item を返す
//}
//
//type Item struct {
//	data [128]uint64
//}
//
//func parse(_ [][]byte) {
//	// 実際は CSV 解析や I/O など
//}

/*----------------- benchmark ------------------*/
func main() {
	var total float64

	for i := 1; i <= runs; i++ {
		start := time.Now()
		res := buildLargeResponse()
		printMemUsage()
		sec := time.Since(start).Seconds()
		total += sec
		fmt.Printf("Run %3d: %.6f seconds (%s)\n", i, sec, res)
	}
	fmt.Printf("Average over %d runs: %.6f seconds\n",
		runs, total/float64(runs))
}
