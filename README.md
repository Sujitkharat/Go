# Golang Interview Preparation Guide

This README file contains categorized Golang interview questions with detailed explanations and code examples where necessary.

---

## ✅ 1. Language Fundamentals

### 1. What are the differences between `var`, `:=`, and `const`?

- `var`: Used to declare a variable with a type (explicit or inferred).
  ```go
  var x int = 5
  var y = "hello"
  ```
- `:=`: Short-hand for declaring and initializing a variable (inside functions only).
  ```go
  z := 10
  ```
- `const`: Used to declare constants. These values cannot be changed after declaration.
  ```go
  const Pi = 3.14
  ```

---

### 2. How does Go handle memory allocation and garbage collection?

Go uses automatic memory management with a garbage collector. Variables are allocated either on the stack or the heap, depending on their usage. Unreferenced memory is automatically reclaimed.

---

### 3. What are slices and how are they different from arrays?

- Arrays are fixed-size.
  ```go
  var arr [3]int = [3]int{1, 2, 3}
  ```
- Slices are dynamic and backed by arrays.
  ```go
  s := []int{1, 2, 3}
  ```

Slices can grow using `append()`. Internally, they point to an array and include length and capacity.

---

### 4. How are maps implemented in Go?

Maps in Go are hash tables. You declare and initialize them like:
```go
m := make(map[string]int)
m["one"] = 1
```
They are not thread-safe, and concurrent access requires synchronization.

---

### 5. What are the zero values of different data types in Go?

| Type       | Zero Value  |
|------------|-------------|
| int        | 0           |
| string     | ""          |
| bool       | false       |
| pointer    | nil         |
| slice/map  | nil         |
| struct     | All fields zero-initialized |

---

### 6. How is error handling done in Go?

Errors are returned as the last value from a function:
```go
func divide(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("division by zero")
  }
  return a / b, nil
}
```

---

### 7. What is a struct? How do you define and use methods on a struct?

```go
type Person struct {
  Name string
  Age  int
}

func (p Person) Greet() string {
  return "Hello, " + p.Name
}
```

---

### 8. How do interfaces work in Go?

Interfaces define method sets. Any type implementing those methods satisfies the interface.
```go
type Speaker interface {
  Speak() string
}
```
Go uses "duck typing": if it quacks like a duck...

---

### 9. What is the difference between a pointer and a value receiver?

Pointer receivers allow you to modify the actual struct. Value receivers operate on a copy.

---

### 10. What does `defer` do?

It defers a function's execution until the surrounding function returns. Useful for cleanup.

---

## ⚙️ 2. Concurrency

### 1. What are goroutines?

Lightweight threads managed by Go runtime:
```go
go func() {
  fmt.Println("Hello from goroutine")
}()
```

---

### 2. How does the `select` statement work?

It lets you wait on multiple channel operations:
```go
select {
case msg := <-ch:
  fmt.Println(msg)
case <-time.After(time.Second):
  fmt.Println("timeout")
}
```

---

### 3. What are channels?

They allow goroutines to communicate:
```go
ch := make(chan int)
ch <- 42
x := <-ch
```

---

### 4. Buffered vs. Unbuffered Channels

- Buffered:
  ```go
  ch := make(chan int, 2)
  ```
- Unbuffered: sends block until receive happens.

---

### 5. Race Conditions

Occurs when goroutines access shared data without synchronization. Use `sync.Mutex`.

---

### 6. sync.WaitGroup

Waits for a group of goroutines to finish:
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
  defer wg.Done()
}()
wg.Wait()
```

---

### 7. sync.Mutex and sync.RWMutex

- `Mutex`: mutual exclusion
- `RWMutex`: allows multiple readers or one writer

---

## 🧰 3. Tools, Packages, and Testing

### 1. Dependency Management (`go mod`)
```bash
go mod init mymodule
go get github.com/pkg/errors
```

---

### 2. Unit Tests
```go
func TestAdd(t *testing.T) {
  if Add(2, 2) != 4 {
    t.Fail()
  }
}
```

---

### 3. Benchmark Tests
```go
func BenchmarkAdd(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Add(1, 2)
  }
}
```

---

### 4. Build Tags
Use `// +build` directives to conditionally include code.

---

### 5. Profiling Tools
Use `pprof`:
```go
import _ "net/http/pprof"
```

---

## 🚀 4. Advanced Topics

### 1. Garbage Collection

Uses concurrent mark-and-sweep collector. Helps minimize pause times.

---

### 2. Performance Optimization

Avoid allocations in hot paths, reuse memory with sync.Pool.

---

### 3. Global Variables and Concurrency

Guard with `sync.Mutex` or use atomic operations.

---

### 4. Context Package
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

---

### 5. Graceful Shutdown
Use `context` and OS signal notifications.

---

## 🧩 5. Best Practices

### 1. Avoid panic/recover

Use it only for truly unexpected conditions.

---

### 2. Large Structures

Pass pointers to avoid copying.

---

### 3. Copying a Mutex

Do not do this. Causes undefined behavior.

---

### 4. new vs make

- `new`: allocates memory, returns pointer
- `make`: for slices, maps, channels

---

### 5. Interface vs Concrete Type

Use interfaces when you want abstraction and flexibility.

---

## 🌐 6. Web & API Development

### 1. Basic REST API
```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, World!")
})
```

---

### 2. Middleware
Wrap handlers for logging, authentication, etc.

---

### 3. Timeouts
Use `http.Server{ReadTimeout: ..., WriteTimeout: ...}`

---

End of README
