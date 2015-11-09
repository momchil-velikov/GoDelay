// Copyright (C) 2015 Momchil Velikov. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package delay

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func atomicInc(p *uint32) uint32 {
	return atomic.AddUint32(p, 1) - 1
}

func TestBasicTimeouts(t *testing.T) {
	done := []int{0, 0, 0, 0, 0}
	cnt := uint32(0)
	cl := New()
	cl.Schedule(200*time.Millisecond, func() { done[atomicInc(&cnt)] = 2 })
	cl.Schedule(100*time.Millisecond, func() { done[atomicInc(&cnt)] = 1 })
	cl.Schedule(500*time.Millisecond, func() { done[atomicInc(&cnt)] = 5 })
	cl.Schedule(300*time.Millisecond, func() { done[atomicInc(&cnt)] = 3 })
	cl.Schedule(400*time.Millisecond, func() { done[atomicInc(&cnt)] = 4 })
	time.Sleep(2 * time.Second)
	cl.Stop()
	if done[0] != 1 || done[1] != 2 || done[2] != 3 || done[3] != 4 || done[4] != 5 {
		t.Error("invalid timeout order")
	}
}

func TestZillionTimeouts(t *testing.T) {
	done := uint32(0)
	cl := New()
	const N = 1000000
	for i := 0; i < N; i++ {
		d := time.Duration(rand.Intn(400))
		cl.Schedule(d*time.Millisecond, func() { atomicInc(&done) })
	}
	time.Sleep(2 * time.Second)
	cl.Stop()
	if done != N {
		t.Errorf("invalid callback count: expected %d, got %d", N, done)
	}
}
