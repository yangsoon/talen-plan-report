package main

import (
	"runtime"
	"strconv"
	"unsafe"

	"github.com/pingcap/tidb/util/mvmap"
)

// Join accepts a join query of two relations, and returns the sum of
// relation0.col0 in the final result.
// Input arguments:
//   f0: file name of the given relation0
//   f1: file name of the given relation1
//   offset0: offsets of which columns the given relation0 should be joined
//   offset1: offsets of which columns the given relation1 should be joined
// Output arguments:
//   sum: sum of relation0.col0 in the final result
func JoinBase(f0, f1 string, offset0, offset1 []int) (sum uint64) {
	tbl0, tbl1 := readCSVFileIntoTbl(f0), readCSVFileIntoTbl(f1)
	flag := false

	if len(tbl0) > len(tbl1) {
		tbl0, tbl1 = tbl1, tbl0
		offset0, offset1 = offset1, offset0
		flag = true
	}

	hashtable := buildHashTblBase(tbl0, offset0, flag)
	numCPU := runtime.NumCPU()
	resultCh := make(chan uint64, numCPU)
	batch := len(tbl1) / numCPU
	for i := 0; i < numCPU; i++ {
		start := i * batch
		end := start + batch
		if i == numCPU-1 {
			end = len(tbl1)
		}
		go func() {
			s, e := start, end
			joinBaseWorker(hashtable, tbl1[s:e], tbl0, offset1, resultCh, flag)
		}()
	}
	for i := 0; i < numCPU; i++ {
		sum += <-resultCh
	}
	return
}

func joinBaseWorker(hashtable *mvmap.MVMap, outerSlice [][]string, innertbl [][]string, offset []int, resultCh chan uint64, flag bool) {
	var sum uint64

	var keyHash []byte
	var vals [][]byte
	for _, row := range outerSlice {
		for i, off := range offset {
			if i > 0 {
				keyHash = append(keyHash, '_')
			}
			keyHash = append(keyHash, []byte(row[off])...)
		}
		vals = hashtable.Get(keyHash, vals)
		keyHash = keyHash[:0]
		switch flag {
		case true:
			// 当使用右表进行probe的时候，我们取出左表对应的row[0]值和从hash表中查找到的数组长度相乘，用来计算sum。
			v, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				panic("joinBaseWorker Convert\n" + err.Error())
			}
			t := v * int64(len(vals))
			sum += uint64(t)
		case false:
			// 当使用左表进行probe的时候，我们从hash表获取到相应的值之后进行累加进行就可以了。
			for _, val := range vals {
				v := *(*int64)(unsafe.Pointer(&val[0]))
				sum += uint64(v)
			}
		}
		vals = vals[:0]
	}
	resultCh <- sum
}

func buildHashTblBase(tbl [][]string, offset []int, flag bool) (hashtable *mvmap.MVMap) {
	var keyBuffer []byte
	var valBuffer []byte
	hashtable = mvmap.NewMVMap()
	if flag {
		valBuffer = make([]byte, 1)
	} else {
		valBuffer = make([]byte, 8)
	}
	for _, row := range tbl {
		for j, off := range offset {
			if j > 0 {
				keyBuffer = append(keyBuffer, '_')
			}
			keyBuffer = append(keyBuffer, []byte(row[off])...)
		}
		switch flag {
		case true:
			// 当使用右表进行hash的时候，我们只要简单的向keyBuffer对应的值中插入一个大小为0的[]byte数组就行，因为这里值是用来存储数据长度的。
			hashtable.Put(keyBuffer, valBuffer)
		case false:
			// 当使用左表进行hash的时候，我们这里选择存储row[0]的值，在probe阶段就不用再去查找左表来计算sum值了。
			v, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				panic("hashWorker Convert\n" + err.Error())
			}
			*(*int64)(unsafe.Pointer(&valBuffer[0])) = int64(v)
			hashtable.Put(keyBuffer, valBuffer)
		}
		keyBuffer = keyBuffer[:0]
	}
	return
}
