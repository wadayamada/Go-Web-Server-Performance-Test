package main

import (
	"fmt"
	"time"
)

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
func buildItems() uint64 {
	var total uint64
	for i := 0; i < chunkCnt; i++ {
		// Item は値として生成→ポインタを外に出さないので escape しない
		it := createItemAddress()
		total += it.data[0] // ダミー集計
	}
	return total
}

func createItem() Item {
	it := Item{}
	// 適当に 0..127 を書き込む
	for j := 0; j < len(it.data); j++ {
		it.data[j] = uint64(j)
	}
	return it // スタック上の Item を返す
}

func createItemAddress() *Item {
	it := Item{}
	// 適当に 0..127 を書き込む
	for j := 0; j < len(it.data); j++ {
		it.data[j] = uint64(j)
	}
	return &it // スタック上の Item を返す
}

type Item struct {
	data [128]uint64
}

/*----------------- benchmark ------------------*/
func main() {
	var total float64
	// レスポンス速度を入れる配列
	resps := make([]float64, 0, runs)

	for i := 1; i <= runs; i++ {
		start := time.Now()
		res := buildItems()
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
