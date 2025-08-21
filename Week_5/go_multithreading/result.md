

Mutex len(m) = 50000, Time = 2.159945ms (~2 ms)
Mutex len(m) = 50000, Time = 1.878116ms (~1 ms)
sync.Map size=50000 time=25.704937ms (~25 ms)

Unbuffered: 148.978834ms (~148 ms)
Buffered:   4.192083ms (~4 ms)


// context switch
Single-thread   total: 188.997375ms, avg switch: 94ns
Multi-thread    total: 257.260208ms, avg switch: 128ns