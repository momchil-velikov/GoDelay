// Copyright (C) 2015 Momchil Velikov. All Rights Reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package delay allows scheduling of functions to be invoked
// after a timeout. It provides an alternative to the idiomatic Go way of
// doing
//   go func() { time.Sleep(timeout); doStuff() }()
// when there's a need to schedule thousands or millions
// of timeouts.
package delay

import (
	"time"
)

// Type of the delayed call facility instances.
type DelayCaller struct {
	queue   callQueue
	queueIn chan call
}

// Create a new delayed call facility instance.
func New() *DelayCaller {
	var c DelayCaller
	c.queue = nil
	c.queueIn = make(chan call)
	go c.runner()
	return &c
}

// This goroutine is the one, that actually executes the functions.
func (cl *DelayCaller) runner() {
	for {
		// Run all the functions in the queue, whose timeout has expired. After
		// the loop, the variable |delay| contains the time until the next timeout.
		delay := time.Duration(0)
		for cl.queue.Len() > 0 {
			c := cl.queue.topCall()
			now := time.Now()
			if now.Before(c.Tm) {
				delay = c.Tm.Sub(now)
				break
			}
			cl.queue.popCall()
			c.Fn()
		}
		// Wait for a new timeout request, but block at most until the next
		// already scheduled timeout.
		var timeout <-chan time.Time
		if delay == 0 {
			timeout = nil
		} else {
			timeout = time.After(delay)
		}
		select {
		case c, ok := <-cl.queueIn:
			if !ok {
				return
			} else {
				cl.queue.pushCall(c)
			}
		case <-timeout:
		}
	}
}

// Schedule a function to be called after the specified timeout.
func (cl *DelayCaller) Schedule(d time.Duration, fn func()) {
	tm := time.Now().Add(d)
	cl.queueIn <- call{fn, tm}
}

// Stop the facility.
func (cl *DelayCaller) Stop() {
	close(cl.queueIn)
}
