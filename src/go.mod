module src

require (
	github.com/Code-Hex/go-generics-cache v1.0.0
	policy/clock v0.0.1
	policy/fifo v0.0.1
	policy/lfu v0.0.1
	policy/lru v0.0.1
	policy/mru v0.0.1
)

require golang.org/x/exp v0.0.0-20220209042442-160e291fcf24 // indirect

replace (
	policy/clock => ./../../../GolandProjects/DistributedCacheSimulator/src/policy/clock
	policy/fifo => ./../../../GolandProjects/DistributedCacheSimulator/src/policy/fifo
	policy/lfu => ./../../../GolandProjects/DistributedCacheSimulator/src/policy/lfu
	policy/lru => ./../../../GolandProjects/DistributedCacheSimulator/src/policy/lru
	policy/mru => ./../../../GolandProjects/DistributedCacheSimulator/src/policy/mru
)

go 1.18
