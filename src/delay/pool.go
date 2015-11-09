package delay

type pool struct {
	n    int           // number of threads
	in   chan func()   // pass functions for execution, close for stop
	down chan struct{} // receive stop response
}

func (p *pool) init(n int) *pool {
	p.n = n
	p.in = make(chan func())
	p.down = make(chan struct{})
	for i := 0; i < n; i++ {
		go p.runner()
	}
	return p
}

func (p *pool) runner() {
	for fn := range p.in {
		fn()
	}
	p.down <- struct{}{}
}

func (p *pool) send(fn func()) {
	p.in <- fn
}

func (p *pool) stop() {
	close(p.in)
	for i := 0; i < p.n; i++ {
		<-p.down
	}
}
