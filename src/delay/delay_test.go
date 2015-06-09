// Copyright (C) 2015 Momchil Velikov. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package delay

import (
	"math/rand"
	"testing"
	"time"
)

func TestBasicTimeouts(t *testing.T) {
	done := []int{0, 0, 0, 0, 0}
	cnt := 0
	cl := New()
	cl.Schedule(200*time.Millisecond, func() { done[cnt] = 2; cnt++ })
	cl.Schedule(100*time.Millisecond, func() { done[cnt] = 1; cnt++ })
	cl.Schedule(500*time.Millisecond, func() { done[cnt] = 5; cnt++ })
	cl.Schedule(300*time.Millisecond, func() { done[cnt] = 31; cnt++ })
	cl.Schedule(300*time.Millisecond, func() { done[cnt] = 32; cnt++ })
	time.Sleep(2 * time.Second)
	cl.Stop()
	if done[0] != 1 || done[1] != 2 || done[2] != 31 || done[3] != 32 || done[4] != 5 {
		t.Error("invalid timeout order")
	}
}

func TestZillionTimeouts(t *testing.T) {
	var done int = 0
	cl := New()
	const N = 1000000
	for i := 0; i < N; i++ {
		d := time.Duration(rand.Intn(400))
		cl.Schedule(d*time.Millisecond, func() { done++ })
	}
	time.Sleep(2 * time.Second)
	cl.Stop()
	if done != N {
		t.Errorf("invalid callback count: expected %d, got %d", N, done)
	}
}
