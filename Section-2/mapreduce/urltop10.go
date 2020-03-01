package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var URLKind = 1024

type URLItem struct {
	url string
	cnt int
}

type URLTopK []URLItem

// URLTop10 .
func URLTop10(nWorkers int) (args RoundsArgs) {

	args = append(args, RoundArgs{
		MapFunc:    URLCountMap,
		ReduceFunc: URLCountReduce,
		NReduce:    nWorkers,
	})

	args = append(args, RoundArgs{
		MapFunc:    TopKMergeMap,
		ReduceFunc: GetTopKReduce,
		NReduce:    1,
	})

	return
}

func URLCountMap(filename string, contents string) (kvs []KeyValue) {
	lines := strings.Split(contents, "\n")
	kv := make(map[string]int, URLKind)

	var kindCount int

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		if _, ok := kv[l]; !ok {
			kindCount++
		}
		kv[l]++
	}

	kvs = make([]KeyValue, 0, kindCount)
	var buffer bytes.Buffer
	for k, v := range kv {
		buffer.WriteString(k)
		buffer.WriteString(" ")
		buffer.WriteString(strconv.Itoa(v))

		kvs = append(kvs, KeyValue{
			Key:   strconv.Itoa(ihash(k) % GetMRCluster().NWorkers()),
			Value: buffer.String(),
		})

		buffer.Reset()
	}
	return
}

func URLCountReduce(key string, values []string) (res string) {
	kv := make(map[string]int, URLKind)

	for _, value := range values {
		if len(value) == 0 {
			continue
		}
		tmp := strings.Split(value, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		kv[tmp[0]] += n
	}

	topk := Top10(kv)

	buf := new(bytes.Buffer)
	for i := 0; i < len(topk); i++ {
		fmt.Fprintf(buf, "%s %d\n", topk[i].url, topk[i].cnt)
	}
	res = buf.String()
	return
}

func TopKMergeMap(filename string, contents string) (kvs []KeyValue) {
	lines := strings.Split(contents, "\n")
	kvs = make([]KeyValue, 0, len(lines))

	var buffer bytes.Buffer
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		tmp := strings.Split(l, " ")

		buffer.WriteString(tmp[0])
		buffer.WriteString(" ")
		buffer.WriteString(tmp[1])

		kvs = append(kvs, KeyValue{
			Key:   "",
			Value: buffer.String(),
		})
		buffer.Reset()
	}
	return
}

func GetTopKReduce(key string, values []string) (res string) {
	ucs := make([]*URLItem, 0, len(values))

	for _, v := range values {
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		ucs = append(ucs, &URLItem{tmp[0], n})
	}

	sort.Slice(ucs, func(i, j int) bool {
		if ucs[i].cnt == ucs[j].cnt {
			return ucs[i].url < ucs[j].url
		}
		return ucs[i].cnt > ucs[j].cnt
	})

	buf := new(bytes.Buffer)
	for i := 0; i < len(ucs); i++ {
		if i == 10 {
			break
		}
		fmt.Fprintf(buf, "%s: %d\n", ucs[i].url, ucs[i].cnt)
	}
	res = buf.String()
	return
}

func Top10(urlKV map[string]int) (topK URLTopK) {

	c := 0
	topK = make(URLTopK, 0, 10)

	var minItem interface{}
	var minVal int

	for url, num := range urlKV {
		c++
		switch {
		case c > 10:
			if num < minVal {
				continue
			}
			heap.Push(&topK, URLItem{url, num})
			heap.Pop(&topK)
			minItem = heap.Pop(&topK)
			minVal = minItem.(URLItem).cnt
			heap.Push(&topK, minItem)
		case c < 10:
			topK = append(topK, URLItem{url, num})
		case c == 10:
			topK = append(topK, URLItem{url, num})
			heap.Init(&topK)
			minItem = heap.Pop(&topK)
			minVal = minItem.(URLItem).cnt
			heap.Push(&topK, minItem)
		}
	}
	return
}

func (u URLTopK) Len() int {
	return len(u)
}

func (u URLTopK) Less(i, j int) bool {
	if u[i].cnt == u[j].cnt {
		return u[i].url > u[j].url
	}
	return u[i].cnt < u[j].cnt
}

func (u URLTopK) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u *URLTopK) Push(a interface{}) {
	item := a.(URLItem)
	*u = append(*u, item)
}

func (u *URLTopK) Pop() interface{} {
	n := len(*u)
	item := (*u)[n-1]
	*u = (*u)[:n-1]
	return item
}
