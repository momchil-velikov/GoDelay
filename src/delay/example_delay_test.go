// Copyright (C) 2015 Momchil Velikov. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package delay_test

import (
	"delay"
	"math/rand"
	"time"
)

const N = 1000000

func doNothing() {}

// Schedule timeouts with the |delay| package
func delayTimeouts() {
	cl := delay.New()
	for i := 0; i < N; i++ {
		d := time.Duration(rand.Intn(400))
		cl.Schedule(d*time.Millisecond, doNothing)
	}
	time.Sleep(10 * time.Second)
	cl.Stop()
}

// Schedule timeouts in idiomatic Go style
func goTmeouts() {
	for i := 0; i < N; i++ {
		d := time.Duration(rand.Intn(400))
		go func(d time.Duration) {
			time.Sleep(d)
			doNothing()
		}(d)
	}
	time.Sleep(10 * time.Second)
}

func Example() {
	delayTimeouts()
	goTmeouts()
}
