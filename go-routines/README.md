Lightweight threads managed by Go runtime for concurrency.

🧠 Internal:
Go multiplexes thousands of goroutines onto a few OS threads using its own scheduler.


```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()  // runs asynchronously
    time.Sleep(1 * time.Second) // give goroutine time to run
}


```

Goroutine Lifecycle
- Spawn: Using go keyword

- Execute: Scheduled by Go runtime

- Block: On I/O, channels, select, etc.

- Die: When function returns


✅ 4. Common Pitfall – Main Exits Early
```go
go fmt.Println("Hello") // Main may exit before this runs
```
✅ Solution: Use sync.WaitGroup or channel to wait.


### WaitGroup – Waiting for Goroutines

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i) // Pass i as argument to avoid closure issue
	}
	wg.Wait()
}


```
### Channel

```go
ch := make(chan int)

go func() {
    ch <- 10
}()

val := <-ch
fmt.Println(val)

```

### Pros & Cons
🔹 Kyu chahiye?
- For running code concurrently.

- Efficient resource usage (compared to OS threads).

🔹 Pros:
- Extremely cheap (KBs vs MBs).

- Scalable.

🔹 Cons:
- Hard to debug race conditions.

- Unmanaged goroutines can leak memory.

🔹 Industry Use:
- API handlers, background jobs, real-time systems.

