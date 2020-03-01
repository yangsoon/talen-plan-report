package main

import (
	"runtime"
	"sort"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func merge1(input ...<-chan int64) <-chan int64 {
	if len(input) == 1 {
		return input[0]
	}
	mid := len(input) / 2
	return subMerge1(merge1(input[:mid]...), merge1(input[mid:]...))
}

func subMerge1(input1, input2 <-chan int64) <-chan int64 {
	out := make(chan int64)

	go func() {
		item1, ok1 := <-input1
		item2, ok2 := <-input2

		for ok1 || ok2 {
			if !ok1 || (ok2 && item2 <= item1) {
				out <- item2
				item2, ok2 = <-input2
			} else {
				out <- item1
				item1, ok1 = <-input1
			}
		}
		close(out)
	}()

	return out
}

func arr2chan(arr []int64) <-chan int64 {
	out := make(chan int64, 1024)
	go func() {
		for _, n := range arr {
			out <- n
		}
		close(out)
	}()
	return out
}

func collectResult(src []int64, output <-chan int64) {
	for i := 0; i < len(src); i++ {
		src[i] = <-output
	}
}

func BadMergeSort(src []int64) {

	// 返回逻辑cpu核数，分配多个任务
	number := runtime.NumCPU()

	if len(src) < number {
		number = len(src)
	}

	arrs := make(chan []int64, number)
	batch := len(src) / number

	var wg sync.WaitGroup
	wg.Add(number)

	// 将数据分发到不同的核上执行
	for i := 0; i < number; i++ {
		go func(index int) {
			defer wg.Done()

			size := batch
			offset := index * batch

			if index == number-1 {
				size = len(src) - offset
			}

			cpy := make([]int64, size)
			copy(cpy, src[offset:offset+size])

			sort.Slice(cpy, func(i, j int) bool {
				return cpy[i] < cpy[j]
			})

			arrs <- cpy
		}(i)
	}

	// 等待结果
	wg.Wait()
	close(arrs)

	var inputs []<-chan int64

	for arr := range arrs {
		inputs = append(inputs, arr2chan(arr))
	}

	output := merge1(inputs...)

	collectResult(src, output)
}
