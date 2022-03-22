module src

require (
	github.com/Code-Hex/go-generics-cache v1.0.0
	src/policy/clock v0.0.1
	src/policy/fifo v0.0.1
	src/policy/lfu v0.0.1
	src/policy/lru v0.0.1
	src/policy/mru v0.0.1
)

require golang.org/x/exp v0.0.0-20220209042442-160e291fcf24 // indirect

replace (
	src/policy/clock => ./policy/clock
	src/policy/fifo => ./policy/fifo
	src/policy/lfu => ./policy/lfu
	src/policy/lru => ./policy/lru
	src/policy/mru => ./policy/mru
)

go 1.18
