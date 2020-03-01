package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/pingcap/check"
)

var _ = check.Suite(&sortTestSuite{})

func TestT(t *testing.T) {
	check.TestingT(t)
}

func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int63()
	}
}

type sortTestSuite struct{}

func (s *sortTestSuite) TestMergeSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		MergeSort(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) TestMerge(c *check.C) {
	lens := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		interSrc = make([]int64, lens[i])
		mid := lens[i] / 2
		srcLeftSlice := src[:mid]
		sort.Slice(srcLeftSlice, func(i, j int) bool { return srcLeftSlice[i] < srcLeftSlice[j] })
		srcRightSlice := src[mid:]
		sort.Slice(srcRightSlice, func(i, j int) bool { return srcRightSlice[i] < srcRightSlice[j] })
		merge(src, 0, mid, lens[i])
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) TestCoreSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		interSrc = make([]int64, lens[i])
		coreSort(src, 0, lens[i])
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) TestB2UpMerge(c *check.C) {
	cpus := []int{2, 4, 8, 16, 32, 64, 128}
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		interSrc = make([]int64, lens[i])
		for j := range cpus {
			prepare(src)
			copy(expect, src)
			sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
			numCPU := cpus[j]
			batch := lens[i] / numCPU
			parts := make([]partSrc, numCPU)
			for k := 0; k < numCPU; k++ {
				start := k * batch
				end := start + batch
				if k == numCPU-1 {
					end = lens[i]
				}
				parts[k] = partSrc{start, end}
				srcSlice := src[start:end]
				sort.Slice(srcSlice, func(i, j int) bool { return srcSlice[i] < srcSlice[j] })
			}
			b2UpMerge(src, parts)
			for i := 0; i < len(src); i++ {
				c.Assert(src[i], check.Equals, expect[i])
			}
		}
	}
}
