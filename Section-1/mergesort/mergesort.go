package main

import (
	"runtime"
	"sort"
	"sync"
)

var interSrc []int64

type partSrc struct {
	start int
	end   int
}

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	length := len(src)
	numCPU := runtime.NumCPU()

	if length < numCPU {
		sort.Slice(src, func(i, j int) bool {
			return src[i] < src[j]
		})
		return
	}

	interSrc = make([]int64, length)
	batch := length / numCPU
	parts := make([]partSrc, numCPU)
	var wg sync.WaitGroup
	wg.Add(numCPU)

	for i := 0; i < numCPU; i++ {
		start := i * batch
		end := start + batch
		if i == numCPU-1 {
			end = length
		}
		parts[i] = partSrc{start, end}
		go func(start, end int) {
			defer wg.Done()
			coreSort(src, start, end)
		}(start, end)
	}

	wg.Wait()
	b2UpMerge(src, parts)
}

func merge(src []int64, start, mid, end int) {
	left := start
	right := mid
	idx := start
	for left < mid && right < end {
		if src[left] > src[right] {
			interSrc[idx] = src[right]
			right++
		} else {
			interSrc[idx] = src[left]
			left++
		}
		idx++
	}

	for left < mid {
		interSrc[idx] = src[left]
		left++
		idx++
	}

	for right < end {
		interSrc[idx] = src[right]
		right++
		idx++
	}

	for i := start; i < end; i++ {
		src[i] = interSrc[i]
	}
}

func coreSort(src []int64, start, end int) {
	if end-start <= 1 {
		return
	}
	mid := start + (end-start)>>1
	coreSort(src, start, mid)
	coreSort(src, mid, end)
	merge(src, start, mid, end)
}

func b2UpMerge(src []int64, parts []partSrc) {
	n := len(parts)
	for size := 1; size < n; size *= 2 {
		var wg sync.WaitGroup

		for low := 0; low < n-size; low += size * 2 {
			start := parts[low].start
			mid := parts[low+size-1].end
			endIdx := low + size*2 - 1
			if endIdx > n-1 {
				endIdx = n - 1
			}
			end := parts[endIdx].end
			wg.Add(1)
			go func(start, mid, end int) {
				defer wg.Done()
				merge(src, start, mid, end)
			}(start, mid, end)
		}
		wg.Wait()
	}
}
