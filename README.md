# delay
--
    import "delay"

Package delay allows scheduling of functions to be invoked after a timeout. It
provides an alternative to the idiomatic Go way of doing

    go func() { time.Sleep(timeout); doStuff() }()

when there's a need to schedule *thousands or millions of timeouts*.

## Usage

#### type DelayCaller

```go
type DelayCaller struct {
    // contains filtered or unexported fields
}
```

Type of the delayed call facility instances.

#### func  New

```go
func New() *DelayCaller
```
Create a new delayed call facility instance.

#### func (*DelayCaller) Schedule

```go
func (cl *DelayCaller) Schedule(d time.Duration, fn func())
```
Schedule a function to be called after the specified timeout.

#### func (*DelayCaller) Stop

```go
func (cl *DelayCaller) Stop()
```
Stop the facility.
