Pointers store memory address of a variable.

```go
func change(val *int) {
    *val = 100
}
```

🔹 Kyu chahiye?
- Pass by reference instead of value.

- Avoid copying large structs.

- Shared mutable state.

🔹 Pros:
- Efficient memory usage.

- Useful in performance-critical code.

🔹 Cons:
- Pointer misuse = bugs, nil panic.

- Can lead to harder debugging.

🔹 Industry Use:
- Working with large data structures (e.g., HTTP request structs).

- Sharing config or state across services.