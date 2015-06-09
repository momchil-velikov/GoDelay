// Copyright (C) 2015 Momchil Velikov. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package delay

import (
	"container/heap"
	"time"
)

// The callQueue type implements a priority queue of delayed calls,
// ordered by timeout (absolute time value).
type call struct {
	Fn func()
	Tm time.Time
}

type callQueue []call

func (pq callQueue) Len() int { return len(pq) }

func (pq callQueue) Less(i, j int) bool {
	return pq[i].Tm.Before(pq[j].Tm)
}

func (pq callQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *callQueue) Push(x interface{}) {
	cl := x.(call)
	*pq = append(*pq, cl)
}

func (pq *callQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	cl := old[n-1]
	*pq = old[0 : n-1]
	return cl
}

func (pq *callQueue) pushCall(cl call) {
	heap.Push(pq, cl)
}

func (pq *callQueue) popCall() {
	heap.Pop(pq)
}

func (pq callQueue) topCall() call {
	return pq[0]
}
